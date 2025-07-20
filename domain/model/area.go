package model

import "time"

type Area struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CityID    *uint     `gorm:"column:city_id" json:"city_id"`
	Type      string    `gorm:"type:varchar(10);not null;default:zone" json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
