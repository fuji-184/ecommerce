package models

type Detail_Keranjang struct {
  Id int `json:"id"`
  Id_keranjang int `json:"id_keranjang"`
  Id_item int `json:"id_item"`
  Jumlah int `json:"jumlah"`
  Total_harga float64 `json:"total_harga"`
}