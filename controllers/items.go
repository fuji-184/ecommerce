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

func QueryAllItems(pool *pgxpool.Pool) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request){
    
    rows, err := pool.Query(context.Background(), "select * from items")
    
    if err != nil {
      panic("error saat mengambil data items" + err.Error())
    }
    
    defer rows.Close()
    
    var items []models.Items
    for rows.Next(){
      var item models.Items
      err := rows.Scan(&item.Id, &item.Nama, &item.Harga, &item.Stok, &item.Gambar)
      if err != nil {
        panic("error saat menscan hasil query users " + err.Error())
      }
      
      items = append(items, item)
    }
    
    w.Header().Set("content-type", "application/json")
    w.WriteHeader(http.StatusOK)
    
    json.NewEncoder(w).Encode(items)
  }
}

func QueryOneItem(pool *pgxpool.Pool) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request){
    id := chi.URLParam(r, "id")
    
    rows, err := pool.Query(context.Background(), "select * from items where id = $1", id)
    if err != nil {
      panic("error saat mengambil data: " + err.Error())
    }
    
    defer rows.Close()
    
    var items []models.Items
      for rows.Next(){
        var item models.Items
        err := rows.Scan(&item.Id, &item.Nama, &item.Harga, &item.Stok, &item.Gambar)
        if err != nil {
          panic("error saat menscan hasil query users " + err.Error())
        }
        
        fmt.Println(item.Nama)
        
        items = append(items, item)
      }
      
      w.Header().Set("content-type", "application/json")
      w.WriteHeader(http.StatusOK)
      json.NewEncoder(w).Encode(items)
  }
}

func UpdateItem(pool *pgxpool.Pool) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request){
    var data models.Items
    
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
      panic("error saat mendecode json: " + err.Error())
    }
    
    _, err = pool.Query(context.Background(), "update items set nama = $1, harga = $2, stok = $3 where id = $4;", data.Nama, data.Harga, data.Stok, data.Id)
    if err != nil {
      panic("error saat mengupdate data items: " + err.Error())
    }
  }
}

func AddItem(pool *pgxpool.Pool) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request){
    var item models.Items
    
    err := json.NewDecoder(r.Body).Decode(&item)
    if err != nil {
      panic("error saat mendecode json item: " + err.Error())
    }
    
    _, err = pool.Query(context.Background(), "insert into items (nama, harga, stok) values ($1, $2, $3)", item.Nama, item.Harga, item.Stok)
    if err != nil {
      panic("error saat memasukkan data items: " + err.Error())
    }
  }
}

func DeleteItem(pool *pgxpool.Pool) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    
    _, err := pool.Query(context.Background(), "delete from items where id = $1", id)
    if err != nil {
      panic("error saat menghapus data item: " + err.Error())
    }
    
  }
}