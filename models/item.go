package models

type Items struct {
  Id int `json:"id"`
  Nama string `json:"nama"`
  Harga float64 `json:"harga"`
  Stok int `json:"stok"`
  Gambar string `json:"gambar"`
}