package model

import "time"

type Lumiere struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AreaID    uint      `gorm:"not null" json:"area_id"`
	Power     bool      `gorm:"default:false" json:"power"`
	Dim       int       `gorm:"check:dim >= 0 AND dim <= 100;default:0" json:"dim"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Area      Area      `gorm:"foreignKey:AreaID" json:"-"`
}
