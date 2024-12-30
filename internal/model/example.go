package model

type Example struct {
	ID   int    `column:"id AUTO_INCREMENT" json:"id"`
	Name string `column:"name" json:"name"`
}
