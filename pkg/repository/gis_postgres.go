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
	query := "SELECT *, ST_AsGeoJSON(geom) as geom_json, tlu.tlu_name, cult.title, cult2.title as title2 FROM field_param LEFT JOIN tlu ON field_param.tlu = tlu.id LEFT JOIN cult ON field_param.crop = cult.id LEFT JOIN cult as cult2 ON field_param.crop = cult2.id"
	err := r.db.Select(&fields, query)

	if err != nil {
		return gogis.GeoJSON{}, err
	}

	var features []gogis.Feature
	for _, field := range fields {
		properties := map[string]interface{}{
			"Id_eu":       field.Id_eu,
			"Farm_name":   field.Farm_name,
			"Tlu":         field.Tlu,
			"Tlu_name":    field.Tlu_name,
			"Area_f":      field.Area_f,
			"Organic":     field.Organic,
			"El_p":        field.El_p,
			"Class_p":     field.Class_p,
			"El_k":        field.El_k,
			"Class_k":     field.Class_k,
			"Humus_class": field.Humus_class,
			"ph_water":    field.Ph_water,
			"Date":        field.Date,
			"Mlrt_d":      field.Mlrt_d,
			"Mlrt_w":      field.Mlrt_w,
			"fact_isp":    field.Fact_isp,
			"Crop":        field.Crop,
			"Crop_old":    field.Crop_old,
			"S_type":      field.S_type,
			"S_ptype":     field.S_ptype,
			"S_ms":        field.S_ms,
			"S_ind":       field.S_ind,
			"El_ca":       field.El_ca,
			"El_mg":       field.El_mo,
			"El_n":        field.El_n,
			"El_co":       field.El_co,
			"El_zn":       field.El_zn,
			"El_cu":       field.El_cu,
			"El_mo":       field.El_mo,
			"El_mn":       field.El_mn,
			"El_s":        field.El_s,
			"El_b":        field.El_b,
			"Il":          field.Il,
			"Glina":       field.Glina,
			"Torf":        field.Torf,
			"Pereyv":      field.Pereyv,
			"Plotn":       field.Plotn,
			"Agreg":       field.Agreg,
			"Field_vl":    field.Field_vl,
			"Nas_osn":     field.Nas_osn,
			"Dgrd_st":     field.DGRD_ST,
			"Yklon":       field.Yklon,
			"Metals":      field.Metals,
			"Met_pb":      field.Met_pb,
			"Met_cd":      field.Met_cd,
			"Met_hg":      field.Met_hg,
			"Met_as":      field.Met_as,
			"Plt_tox":     field.Plt_tox,
			"Plt_oil":     field.Plt_oil,
			"Cs137":       field.Cs137,
			"Sr90":        field.Sr90,
			"Gamma":       field.Gamma,
			"Salinity":    field.Salinity,
			"Solon_na":    field.Solon_Na,
			"Solon_pr":    field.Solon_pr,
			"Wind_er":     field.Wind_er,
			"Water_er":    field.Water_er,
			"Zakust":      field.Zakust,
			"Zales":       field.Zales,
			"Zarast":      field.Zarast,
			"Narushen":    field.Narushen,
			"Zastroen":    field.Zastroen,
			"Prochee":     field.Prochee,
			"Mosh_g":      field.Mosh_g,
			"Mosh_m":      field.Mosh_m,
			"Microbio":    field.Microbio,
			"S_cs137":     field.S_cs137,
			"S_sr90":      field.S_sr90,
			"Solon_d":     field.Solon_d,
			"Eco_na":      field.Eco_na,
			"Eko":         field.Eko,
			"gydr_c":      field.Gydr_c,
			"Title":       field.Title,
			"Title2":      field.Title2,
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

func (r *GisPostgres) GetFieldData() ([]gogis.Field, error) {
	var fields []gogis.Field
	query := "SELECT *, ST_AsGeoJSON(geom) as geom_json FROM field_param"
	err := r.db.Select(&fields, query)

	if err != nil {
		return nil, err
	}

	return fields, nil
}

func (r *GisPostgres) GetCult() ([]gogis.Cult, error) {
	var cults []gogis.Cult
	query := "SELECT * FROM cult"
	err := r.db.Select(&cults, query)

	if err != nil {
		return nil, err
	}

	return cults, nil
}

func (r *GisPostgres) CreateCult(cult gogis.Cult) error {
	query := "INSERT INTO cult (id, title) VALUES ($1, $2)"
	_, err := r.db.Exec(query, cult.Id, cult.Title)
	return err
}

func (r *GisPostgres) UpdateCult(cult gogis.Cult) error {
	query := "UPDATE cult SET id= $1 title = $2,  WHERE id = $1"
	_, err := r.db.Exec(query, cult.Id, cult.Title)
	return err
}

func (r *GisPostgres) DeleteCult(id int) error {
	query := "DELETE FROM cult WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
