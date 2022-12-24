package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Image struct {
	Id  string `json:"id" gorm:"primaryKey"`
	Url string `json:"url" gorm:"not null"`
}

// Value Marshal
func (a Image) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *Image) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
