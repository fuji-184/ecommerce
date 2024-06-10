package models

type Users struct {
  Id int `json:"id"`
  Nama string `json:"nama"`
  Username string `json:"username"`
  Password string `json:"password"`
  Is_admin string `json:"is_admin"`
  Saldo float64 `json:"saldo"`
}