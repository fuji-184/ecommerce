package main

import (
	"net/http"
	"io/fs"
	"embed"
	"context"
	"time"
	"path/filepath"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	
	"ecommerce/controllers"
)

//go:embed all:frontend/build
var sveltekit embed.FS
var tokenAuth *jwtauth.JWTAuth

func main() {
  s, err := fs.Sub(sveltekit, "frontend/build")
  
  if err != nil {
    panic(err)
  }
  
  dbpool, err := pgxpool.New(context.Background(), "postgresql://fuji:fuji123@localhost:5432/ecommerce")
  
  if err != nil {
      panic("error menghubungkan ke database: " + err.Error())
  }
  
  defer dbpool.Close()
  
  SveltekitServer := http.FileServer(http.FS(s))
  
  tokenAuth = jwtauth.New("HS256", []byte("secret"), nil, jwt.WithAcceptableSkew(120*time.Second))
  
	r := chi.NewRouter()
	
  r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*","http://", "http://localhost:5173"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300,
  }))
	r.Use(middleware.Logger)
	
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "files"))
	FileServer(r, "/files", filesDir)

  r.Handle("/", SveltekitServer)
  r.Handle("/_app/*", SveltekitServer)
  r.Handle("/favicon.png", SveltekitServer)
  
  r.Post("/daftar", controllers.Daftar(dbpool, tokenAuth))
  r.Post("/masuk", controllers.Masuk(dbpool, tokenAuth))
  
  r.Get("/api/items", controllers.QueryAllItems(dbpool))
  r.Get("/api/items/{id}", controllers.QueryOneItem(dbpool))
  
  r.Route("/api", func(r chi.Router){
    r.Use(jwtauth.Verifier(tokenAuth))
    r.Use(jwtauth.Authenticator(tokenAuth))
    
    r.Get("/users", controllers.QueryAllUsers(dbpool))
    r.Get("/users/{id}", controllers.QueryOne(dbpool))
    r.Put("/users", controllers.UpdateUser(dbpool))
    r.Post("/users", controllers.AddUser(dbpool))
    r.Delete("/users/{id}", controllers.DeleteUser(dbpool))
    
    r.Put("/items", controllers.UpdateItem(dbpool))
    r.Post("/items", controllers.AddItem(dbpool))
    r.Delete("/items/{id}", controllers.DeleteItem(dbpool))
  })
  
  r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request){
    r.URL.Path = "/"
    SveltekitServer.ServeHTTP(w, r)
  })

	http.ListenAndServe("localhost:3000", r)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}