package model

type User struct {
	Id      int    `gorm:"primaryKey" json:"id"`
	Prenom  string `json:"prenom"`
	Nom     string `json:"nom"`
	Adresse string `json:"adresse"`
}
