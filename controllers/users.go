package controllers

import (
  "context"
  "encoding/json"
  "net/http"
  "fmt"
  
  "github.com/go-chi/chi/v5"
  "github.com/jackc/pgx/v5/pgxpool"
  "ecommerce/models"
)

func QueryAllUsers(pool *pgxpool.Pool) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request){
    
    rows, err := pool.Query(context.Background(), "select id, nama, username from users")
    
    if err != nil {
      panic("error saat mengambil data users" + err.Error())
    }
    
    defer rows.Close()
    
    var users []models.Users
    for rows.Next(){
      var user models.Users
      err := rows.Scan(&user.Id, &user.Nama, &user.Username)
      if err != nil {
        panic("error saat menscan hasil query users " + err.Error())
      }
      
      users = append(users, user)
    }
    
    w.Header().Set("content-type", "application/json")
    w.WriteHeader(http.StatusOK)
    
    json.NewEncoder(w).Encode(users)
  }
}

func QueryOne(pool *pgxpool.Pool) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request){
    id := chi.URLParam(r, "id")
    
    rows, err := pool.Query(context.Background(), "select id, nama, username from users where id = $1", id)
    if err != nil {
      panic("error saat mengambil data: " + err.Error())
    }
    
    defer rows.Close()
    
    var users []models.Users
      for rows.Next(){
        var user models.Users
        err := rows.Scan(&user.Id, &user.Nama, &user.Username)
        if err != nil {
          panic("error saat menscan hasil query users " + err.Error())
        }
        
        fmt.Println(user.Nama)
        
        users = append(users, user)
      }
      
      w.Header().Set("content-type", "application/json")
      w.WriteHeader(http.StatusOK)
      json.NewEncoder(w).Encode(users)
  }
}

func UpdateUser(pool *pgxpool.Pool) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request){
    var data models.Users
    
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
      panic("error saat mendecode json: " + err.Error())
    }
    
    _, err = pool.Query(context.Background(), "update users set nama = $1, username = $2 where id = $3;", data.Nama, data.Username, data.Id)
    if err != nil {
      panic("error saat mengupdate data user: " + err.Error())
    }
  }
}

func AddUser(pool *pgxpool.Pool) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request){
    var user models.Users
    
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
      panic("error saat mendecode json user: " + err.Error())
    }
    
    _, err = pool.Query(context.Background(), "insert into users (nama, username, password, is_admin) values ($1, $2, $3, $4)", user.Nama, user.Username, user.Password, user.Is_admin)
    if err != nil {
      panic("error saat memasukkan data user: " + err.Error())
    }
  }
}

func DeleteUser(pool *pgxpool.Pool) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    
    _, err := pool.Query(context.Background(), "delete from users where id = $1", id)
    if err != nil {
      panic("error saat menghapus data user: " + err.Error())
    }
    
  }
}