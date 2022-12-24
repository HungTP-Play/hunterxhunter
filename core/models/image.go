package models

type Image struct {
	Id  string `json:"id" gorm:"primaryKey"`
	Url string `json:"url" gorm:"not null"`
}
