package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ImageUrl     string   `json:"image_url"`
	Nama         string   `json:"nama"`
	Pabrikan     string   `json:"pabrikan"`
	Warna        string   `json:"warna"`
	JenisDupa    string   `json:"jenis_dupa"`
	Aroma        string   `json:"aroma"`
	JenisAroma   string   `json:"jenis_aroma"`
	JenisAbu     string   `json:"jenis_abu"`
	LamaNyala    int      `json:"lama_nyala"`
	Jumlah       int      `json:"jumlah"`
	Panjang      int      `json:"panjang"`
	Berat        int      `json:"berat"`
	Asal         string   `json:"asal"`
	Rating       float64  `json:"rating"`
	JumlahReview int      `json:"jumlah_review"`
	TotalRate    float64  `json:"total_rate"`
	Reviews      []Review `json:"reviews" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
