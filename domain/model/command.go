package model

import (
	"time"

	"gorm.io/datatypes"
)

type Command struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	AreaID       *uint          `json:"area_id"`
	LumiereID    *uint          `json:"lumiere_id"`
	CommandData  datatypes.JSON `gorm:"type:jsonb;not null" json:"command"`
	ScheduledFor *time.Time     `gorm:"column:scheduled_for" json:"scheduled_for"`
	ExecutedAt   *time.Time     `gorm:"column:executed_at" json:"executed_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Area         *Area          `gorm:"foreignKey:AreaID" json:"-"`
	Lumiere      *Lumiere       `gorm:"foreignKey:LumiereID" json:"-"`
}
