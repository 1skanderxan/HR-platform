package domain

import "github.com/google/uuid"

type GeofenceZone struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Latitude  float64   `db:"latitude" json:"latitude"`
	Longitude float64   `db:"longitude" json:"longitude"`
	RadiusM   float64   `db:"radius_m" json:"radius_m"` // metrda
	IsActive  bool      `db:"is_active" json:"is_active"`
}
