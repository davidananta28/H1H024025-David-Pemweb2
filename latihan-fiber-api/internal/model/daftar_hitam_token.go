package model

import "time"

type DaftarHitamToken struct {
	ID              uint      `gorm:"primaryKey"`
	TokenHash       string    `gorm:"size:64;uniqueIndex"`
	KedaluwarsaPada time.Time `gorm:"index"`
	CreatedAt       time.Time
}