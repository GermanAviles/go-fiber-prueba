package model

type Shoe struct {
	// gorm.Model
	Color  string `json:"color"`
	Talla  string `json:"talla"`
	Rating int    `json:"rating"`
}
