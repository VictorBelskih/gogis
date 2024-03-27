package repository

import (
	"encoding/json"

	"github.com/VictorBelskih/gogis"
	"github.com/jmoiron/sqlx"
)

type GisPostgres struct {
	db *sqlx.DB
}

// создание экземпляра репозитория
func NewGisPostgres(db *sqlx.DB) *GisPostgres {
	return &GisPostgres{db: db}
}

func (r *GisPostgres) GetField() (gogis.GeoJSON, error) {
	var fields []gogis.Field
	query := "SELECT Id_eu, Farm_name, Tlu, Area_f, ST_AsGeoJSON(geom) as geom_json FROM field"
	err := r.db.Select(&fields, query)

	if err != nil {
		return gogis.GeoJSON{}, err
	}

	var features []gogis.Feature
	for _, field := range fields {
		properties := map[string]interface{}{
			"Id_eu":     field.Id_eu,
			"Farm_name": field.Farm_name,
			"Tlu":       field.Tlu,
			"Area_f":    field.Area_f,
		}

		feature := gogis.Feature{
			Type:       "Feature",
			Geometry:   json.RawMessage(field.Geom_json),
			Properties: properties,
		}

		features = append(features, feature)
	}

	geoJSON := gogis.GeoJSON{
		Type:     "FeatureCollection",
		Features: features,
	}

	return geoJSON, nil
}
