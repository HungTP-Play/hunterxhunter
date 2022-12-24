package models

type Character struct {
	Id       string  `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Nickname string  `json:"nickname"`
	Images   []Image `json:"images" gorm:"type:jsonb"`
}
