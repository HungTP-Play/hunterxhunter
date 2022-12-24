package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Character struct {
	Id       string  `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Nickname string  `json:"nickname"`
	Images   []Image `json:"images" gorm:"type:jsonb"`
}

// Value Marshal
func (a Character) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *Character) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
