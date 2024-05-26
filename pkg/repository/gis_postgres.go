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
	var fields []gogis.Field // Используйте вашу структуру Field
	query := `SELECT *, ST_AsGeoJSON(geom) as geom_json, tlu.tlu_name, cult.title, cult2.title as title2 FROM field_param
              LEFT JOIN tlu ON field_param.tlu = tlu.id
              LEFT JOIN cult ON field_param.crop = cult.id
              LEFT JOIN cult as cult2 ON field_param.crop = cult2.id`
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
func (r *GisPostgres) GetFieldByUser(id int, role int) (gogis.GeoJSON, error) {
	var fields []gogis.Field // Используйте вашу структуру Field
	query := `SELECT field_param.*, ST_AsGeoJSON(geom) as geom_json, 
	tlu.tlu_name, cult.title, cult2.title as title2, farm.id as id_farm, farm.id_user, kamenistost.name as kamn,
	type_pojv.name as t_pojv, subtype_pojv.name as sub_pojv, mex_sost.name	as mex_sost, narushennye_zemli.name as f_narsh,
	torf.name as f_torf, zagryaz.name as zagryaz, zakustarennost.name as kust, zalesennost.name as les,
	zasorennost.name as sorn, zastroennost.name as stroy, prochie_zemli.name as prch 
	FROM field_param
	LEFT JOIN prochie_zemli ON field_param.Prochee=prochie_zemli.id 
	LEFT JOIN zastroennost ON field_param.Zastroen=zastroennost.id  
	LEFT JOIN zasorennost ON field_param.Zarast=zasorennost.id  
	LEFT JOIN zalesennost ON field_param.Zales=zalesennost.id 
	LEFT JOIN zakustarennost ON field_param.Zakust=zakustarennost.id 
	LEFT JOIN zagryaz ON field_param.Metals=zagryaz.id
	LEFT JOIN torf ON field_param.torf=torf.id
	LEFT JOIN narushennye_zemli ON field_param.Narushen=narushennye_zemli.id
	LEFT JOIN mex_sost ON field_param.S_ms=mex_sost.id
	LEFT JOIN subtype_pojv ON field_param.S_ptype=subtype_pojv.id
	LEFT JOIN kamenistost ON field_param.DGRD_ST=kamenistost.id
	LEFT JOIN type_pojv ON field_param.s_type=type_pojv.id_pojv
	LEFT JOIN tlu ON field_param.tlu = tlu.id
	LEFT JOIN cult ON field_param.crop = cult.id
	LEFT JOIN cult as cult2 ON field_param.crop = cult2.id
	LEFT JOIN farm ON field_param.farm_id = farm.id`

	if role != 1 {
		query += ` WHERE farm.id_user=$1`
	}

	// Добавьте id в качестве аргумента в функцию Select только если пользователь не является админом
	var err error
	if role != 1 {
		err = r.db.Select(&fields, query, id)
	} else {
		err = r.db.Select(&fields, query)
	}

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
			"Kamn":        field.Kamn,
			"T_pojv":      field.T_pojv,
			"Sub_pojv":    field.Sub_pojv,
			"Mex_sost":    field.Mex_sost,
			"F_narsh":     field.F_narsh,
			"F_torf":      field.F_torf,
			"Zagryaz":     field.Zagryaz,
			"Kust":        field.Kust,
			"Les":         field.Les,
			"Sorn":        field.Sorn,
			"Stroy":       field.Stroy,
			"Prch":        field.Prch,
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
func (r *GisPostgres) GetFieldData(id int, role int) ([]gogis.Field, error) {
	var fields []gogis.Field
	query := `SELECT field_param.*, ST_AsGeoJSON(geom) as geom_json, 
	tlu.tlu_name, cult.title, cult2.title as title2, farm.id as id_farm, farm.id_user, kamenistost.name as kamn,
	type_pojv.name as t_pojv, subtype_pojv.name as sub_pojv, mex_sost.name	as mex_sost, narushennye_zemli.name as f_narsh,
	torf.name as f_torf, zagryaz.name as zagryaz, zakustarennost.name as kust, zalesennost.name as les,
	zasorennost.name as sorn, zastroennost.name as stroy, prochie_zemli.name as prch 
	FROM field_param
	LEFT JOIN prochie_zemli ON field_param.Prochee=prochie_zemli.id 
	LEFT JOIN zastroennost ON field_param.Zastroen=zastroennost.id  
	LEFT JOIN zasorennost ON field_param.Zarast=zasorennost.id  
	LEFT JOIN zalesennost ON field_param.Zales=zalesennost.id 
	LEFT JOIN zakustarennost ON field_param.Zakust=zakustarennost.id 
	LEFT JOIN zagryaz ON field_param.Metals=zagryaz.id
	LEFT JOIN torf ON field_param.torf=torf.id
	LEFT JOIN narushennye_zemli ON field_param.Narushen=narushennye_zemli.id
	LEFT JOIN mex_sost ON field_param.S_ms=mex_sost.id
	LEFT JOIN subtype_pojv ON field_param.S_ptype=subtype_pojv.id
	LEFT JOIN kamenistost ON field_param.DGRD_ST=kamenistost.id
	LEFT JOIN type_pojv ON field_param.s_type=type_pojv.id_pojv
	LEFT JOIN tlu ON field_param.tlu = tlu.id
	LEFT JOIN cult ON field_param.crop = cult.id
	LEFT JOIN cult as cult2 ON field_param.crop = cult2.id
	LEFT JOIN farm ON field_param.farm_id = farm.id`

	if role != 1 {
		query += ` WHERE farm.id_user=$1`
	}

	// Добавьте id в качестве аргумента в функцию Select только если пользователь не является админом
	var err error
	if role != 1 {
		err = r.db.Select(&fields, query, id)
	} else {
		err = r.db.Select(&fields, query)
	}

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
func (r *GisPostgres) GetCultByID(id int) (*gogis.Cult, error) {
	var cult gogis.Cult
	query := "SELECT * FROM cult WHERE id = $1"
	err := r.db.Get(&cult, query, id)

	if err != nil {
		return nil, err
	}

	return &cult, nil
}
func (r *GisPostgres) CreateCult(cult gogis.Cult) error {
	query := "INSERT INTO cult (id, title) VALUES ($1, $2)"
	_, err := r.db.Exec(query, cult.Id, cult.Title)
	return err
}

func (r *GisPostgres) UpdateCult(cult gogis.Cult) error {
	query := "UPDATE cult SET id = $2, title = $3 WHERE id = $1"
	_, err := r.db.Exec(query, cult.OldId, cult.Id, cult.Title)
	return err
}

func (r *GisPostgres) DeleteCult(id int) error {
	query := "DELETE FROM cult WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
func (r *GisPostgres) GetDistrict() ([]gogis.District, error) {
	var district []gogis.District
	query := "SELECT * FROM district"
	err := r.db.Select(&district, query)

	if err != nil {
		return nil, err
	}

	return district, nil
}

func (r *GisPostgres) GetFarm() ([]gogis.Farm, error) {
	var farms []gogis.Farm
	query := "SELECT * FROM farm"
	err := r.db.Select(&farms, query)

	if err != nil {
		return nil, err
	}

	return farms, nil
}
func (r *GisPostgres) GetFarmByID(id int) (*gogis.Farm, error) {
	var farm gogis.Farm
	query := "SELECT * FROM farm WHERE id = $1"
	err := r.db.Get(&farm, query, id)

	if err != nil {
		return nil, err
	}

	return &farm, nil
}

func (r *GisPostgres) CreateFarm(farm gogis.Farm) error {
	query := "INSERT INTO farm (id, name, district, id_user) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, farm.Id, farm.Name, farm.District, farm.Id_user)
	return err
}

func (r *GisPostgres) UpdateFarm(farm gogis.Farm) error {
	query := "UPDATE farm SET id = $2, name = $3, district = $4, id_user=$5 WHERE id = $1"
	_, err := r.db.Exec(query, farm.OldId, farm.Id, farm.Name, farm.District, farm.Id_user)
	return err
}

func (r *GisPostgres) DeleteFarm(id int) error {
	query := "DELETE FROM farm WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
