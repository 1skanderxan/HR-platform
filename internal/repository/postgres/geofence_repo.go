package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yourname/hr-portal-backend/internal/domain"
)

type geofenceRepo struct {
	db *sqlx.DB
}

func NewGeofenceRepository(db *sqlx.DB) domain.GeofenceRepository {
	return &geofenceRepo{db: db}
}

func (r *geofenceRepo) GetActiveZone(ctx context.Context) (*domain.GeofenceZone, error) {
	zone := &domain.GeofenceZone{}
	err := r.db.GetContext(ctx, zone,
		`SELECT id, name, latitude, longitude, radius_m, is_active FROM geofence_zones WHERE is_active = true LIMIT 1`)
	if err != nil {
		return nil, err
	}
	return zone, nil
}
