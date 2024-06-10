package controllers

import (
  "github.com/jackc/pgx/v5/pgxpool"
  "net/http"
  "encoding/json"
  "ecommerce/models"
  "context"
  "github.com/go-chi/jwtauth/v5"
  "fmt"
)

func Daftar(pool *pgxpool.Pool, tokenizer *jwtauth.JWTAuth) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request){
    
    var data models.Users
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
      panic(err)
    }
    
    fmt.Println(data.Nama)
    
    var username string
    err = pool.QueryRow(context.Background(), "INSERT INTO users (nama, username, password, is_admin) VALUES ($1, $2, $3, $4) RETURNING username", data.Nama, data.Username, data.Password, "false").Scan(&username)
    if err != nil {
      panic(err)
    }
    
    _, tokenString, _ := tokenizer.Encode(map[string]interface{}{"username": username})
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"nama": data.Nama, "token": tokenString})
        
      }
}

func Masuk(pool *pgxpool.Pool, tokenizer *jwtauth.JWTAuth) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request){
    var user models.Users
    
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
      panic("error saat mendecoding json: " + err.Error())
    }
    
    rows, err := pool.Query(context.Background(), "select nama, username, saldo from users where username = $1 and password = $2", user.Username, user.Password)
    if err != nil {
      panic("error saat mengecek data login: " + err.Error())
    }
    
    if rows.Next() {
      rows.Scan(&user.Nama, &user.Username, &user.Saldo)
    } else {
      http.Error(w, "Username atau Password salah", http.StatusUnauthorized)
      return
    }
    
    _, tokenString, _ := tokenizer.Encode(map[string]interface{}{"username": user.Username})
    
    w.Header().Set("content-type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "pesan": "login berhasil",
        "nama": user.Nama,
        "saldo": user.Saldo,
        "token": tokenString,
    })

  }
}