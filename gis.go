package gogis

type Field struct {
	Id_farm     int     `json:"id_farm "`
	Farm_id     int     `json:"farm_id"`
	F_name      string  `json:"f_name"`
	District    int     `json:"District"`
	Id_user     int     `json:"Id_user"`
	Id          int     `json:"id"`
	Geom        string  `json:"geom"`
	Id_eu       int     `json:"id_eu"`
	Station     string  `json:"station"`
	Date        string  `json:"date"`
	Sub_name    float64 `json:"sub_name"`
	Reg_name    string  `json:"reg_name"`
	Reg_id      string  `json:"reg_id"`
	Reg_oktmo   string  `json:"reg_oktmo"`
	Sp_name     string  `json:"sp_name"`
	Sp_id       string  `json:"sp_id"`
	Date_dzz    string  `json:"date_dzz"`
	Dzz_ka_s    string  `json:"dzz_ka_s"`
	Farm_name   string  `json:"farm_name"`
	Tlu         float64 `json:"tlu"`
	Tlu_name    string  `json:"tlu_name"`
	Mlrt_d      float64 `json:"mlrt_d"`
	Mlrt_w      float64 `json:"mlrt_w"`
	Area_f      float64 `json:"area_f"`
	Crop        int     `json:"crop"`
	Crop_old    int     `json:"crop_old"`
	S_type      float64 `json:"s_type"`
	S_ptype     float64 `json:"s_ptype"`
	S_ms        float64 `json:"s_ms"`
	S_ind       float64 `json:"s_ind"`
	Organic     float64 `json:"organic"`
	Ph_kcl      float64 `json:"ph_kcl"`
	Ph_water    float64 `json:"ph_water"`
	Gydr_c      float64 `json:"gydr_c"`
	Eko         float64 `json:"eko"`
	El_p        float64 `json:"el_p"`
	El_k        float64 `json:"el_k"`
	El_n        float64 `json:"el_n"`
	El_ca       float64 `json:"el_ca"`
	El_mg       float64 `json:"el_mg"`
	El_zn       float64 `json:"el_zn"`
	El_cu       float64 `json:"el_cu"`
	El_mo       float64 `json:"el_mo"`
	El_s        float64 `json:"el_s"`
	El_b        float64 `json:"el_b"`
	Il          float64 `json:"il"`
	Glina       float64 `json:"glina"`
	Torf        float64 `json:"torf"`
	Pereyv      float64 `json:"pereyv"`
	Plotn       float64 `json:"plotn"`
	Agreg       float64 `json:"agreg"`
	Field_vl    float64 `json:"field_vl"`
	Nas_osn     float64 `json:"nas_osn"`
	DGRD_ST     float64 `json:"dgrd_st"`
	Yklon       float64 `json:"yklon"`
	Metals      float64 `json:"metals"`
	Met_pb      float64 `json:"met_pb"`
	Met_cd      float64 `json:"met_cd"`
	Met_hg      float64 `json:"met_hg"`
	Met_as      float64 `json:"met_as"`
	Plt_tox     float64 `json:"plt_tox"`
	Plt_oil     float64 `json:"plt_oil"`
	Cs137       float64 `json:"cs137"`
	Sr90        float64 `json:"sr90"`
	Gamma       float64 `json:"gamma"`
	Salinity    float64 `json:"salinity"`
	Solon_Na    float64 `json:"solon_na"`
	Solon_pr    float64 `json:"solon_pr"`
	Wind_er     float64 `json:"wind_er"`
	Water_er    float64 `json:"water_er"`
	Zakust      float64 `json:"zakust"`
	Zales       float64 `json:"zales"`
	Zarast      float64 `json:"zarast"`
	Zastroen    float64 `json:"zastroen"`
	Narushen    float64 `json:"narushen"`
	Prochee     float64 `json:"prochee"`
	Mosh_g      float64 `json:"mosh_g"`
	Mosh_m      float64 `json:"mosh_m"`
	Microbio    string  `json:"microbio"`
	S_cs137     float64 `json:"s_cs137"`
	S_sr90      float64 `json:"s_sr90"`
	Solon_d     float64 `json:"solon_d"`
	Fact_isp    string  `json:"fact_isp"`
	Eco_na      float64 `json:"eco_na"`
	El_mn       float64 `json:"el_mn"`
	El_co       float64 `json:"el_co"`
	Pol_id      int     `json:"pol_id"`
	Gosprog     int     `json:"gosprog"`
	Rekisp      int     `json:"rekisp"`
	Geom_json   string  `db:"geom_json" json:"geom_json"`
	Humus_class string  `json:"humus_class"`
	Class_k     string  `json:"class_k"`
	Class_p     string  `json:"class_p"`
	Title       string  `json:"title"`
	Title2      string  `json:"Title2"`
}

// package gogis

// import "database/sql"

// type Field struct {
// 	Id_field    int             `json:"id_field"`
// 	Geom        string          `json:"geom"`
// 	Area        float64         `json:"area"`
// 	Farm_id     int             `json:"farm_id"`
// 	F_id        int             `json:"f_id"`
// 	Tlu_field   string          `json:"tlu_field"`
// 	Id          sql.NullInt64   `json:"id"`
// 	Geoms       sql.NullString  `json:"geoms"`
// 	Id_eu       sql.NullInt64   `json:"id_eu"`
// 	Station     sql.NullString  `json:"station"`
// 	Date        sql.NullString  `json:"date"`
// 	Sub_name    sql.NullFloat64 `json:"sub_name"`
// 	Reg_name    sql.NullString  `json:"reg_name"`
// 	Reg_id      sql.NullString  `json:"reg_id"`
// 	Reg_oktmo   sql.NullString  `json:"reg_oktmo"`
// 	Sp_name     sql.NullString  `json:"sp_name"`
// 	Sp_id       sql.NullString  `json:"sp_id"`
// 	Date_dzz    sql.NullString  `json:"date_dzz"`
// 	Dzz_ka_s    sql.NullString  `json:"dzz_ka_s"`
// 	Farm_name   sql.NullString  `json:"farm_name"`
// 	Tlu         sql.NullFloat64 `json:"tlu"`
// 	Tlu_name    sql.NullString  `json:"tlu_name"`
// 	Mlrt_d      sql.NullFloat64 `json:"mlrt_d"`
// 	Mlrt_w      sql.NullFloat64 `json:"mlrt_w"`
// 	Area_f      sql.NullFloat64 `json:"area_f"`
// 	Crop        sql.NullInt64   `json:"crop"`
// 	Crop_old    sql.NullInt64   `json:"crop_old"`
// 	S_type      sql.NullFloat64 `json:"s_type"`
// 	S_ptype     sql.NullFloat64 `json:"s_ptype"`
// 	S_ms        sql.NullFloat64 `json:"s_ms"`
// 	S_ind       sql.NullFloat64 `json:"s_ind"`
// 	Organic     sql.NullFloat64 `json:"organic"`
// 	Ph_kcl      sql.NullFloat64 `json:"ph_kcl"`
// 	Ph_water    sql.NullFloat64 `json:"ph_water"`
// 	Gydr_c      sql.NullFloat64 `json:"gydr_c"`
// 	Eko         sql.NullFloat64 `json:"eko"`
// 	El_p        sql.NullFloat64 `json:"el_p"`
// 	El_k        sql.NullFloat64 `json:"el_k"`
// 	El_n        sql.NullFloat64 `json:"el_n"`
// 	El_ca       sql.NullFloat64 `json:"el_ca"`
// 	El_mg       sql.NullFloat64 `json:"el_mg"`
// 	El_zn       sql.NullFloat64 `json:"el_zn"`
// 	El_cu       sql.NullFloat64 `json:"el_cu"`
// 	El_mo       sql.NullFloat64 `json:"el_mo"`
// 	El_s        sql.NullFloat64 `json:"el_s"`
// 	El_b        sql.NullFloat64 `json:"el_b"`
// 	Il          sql.NullFloat64 `json:"il"`
// 	Glina       sql.NullFloat64 `json:"glina"`
// 	Torf        sql.NullFloat64 `json:"torf"`
// 	Pereyv      sql.NullFloat64 `json:"pereyv"`
// 	Plotn       sql.NullFloat64 `json:"plotn"`
// 	Agreg       sql.NullFloat64 `json:"agreg"`
// 	Field_vl    sql.NullFloat64 `json:"field_vl"`
// 	Nas_osn     sql.NullFloat64 `json:"nas_osn"`
// 	DGRD_ST     sql.NullFloat64 `json:"dgrd_st"`
// 	Yklon       sql.NullFloat64 `json:"yklon"`
// 	Metals      sql.NullFloat64 `json:"metals"`
// 	Met_pb      sql.NullFloat64 `json:"met_pb"`
// 	Met_cd      sql.NullFloat64 `json:"met_cd"`
// 	Met_hg      sql.NullFloat64 `json:"met_hg"`
// 	Met_as      sql.NullFloat64 `json:"met_as"`
// 	Plt_tox     sql.NullFloat64 `json:"plt_tox"`
// 	Plt_oil     sql.NullFloat64 `json:"plt_oil"`
// 	Cs137       sql.NullFloat64 `json:"cs137"`
// 	Sr90        sql.NullFloat64 `json:"sr90"`
// 	Gamma       sql.NullFloat64 `json:"gamma"`
// 	Salinity    sql.NullFloat64 `json:"salinity"`
// 	Solon_Na    sql.NullFloat64 `json:"solon_na"`
// 	Solon_pr    sql.NullFloat64 `json:"solon_pr"`
// 	Wind_er     sql.NullFloat64 `json:"wind_er"`
// 	Water_er    sql.NullFloat64 `json:"water_er"`
// 	Zakust      sql.NullFloat64 `json:"zakust"`
// 	Zales       sql.NullFloat64 `json:"zales"`
// 	Zarast      sql.NullFloat64 `json:"zarast"`
// 	Zastroen    sql.NullFloat64 `json:"zastroen"`
// 	Narushen    sql.NullFloat64 `json:"narushen"`
// 	Prochee     sql.NullFloat64 `json:"prochee"`
// 	Mosh_g      sql.NullFloat64 `json:"mosh_g"`
// 	Mosh_m      sql.NullFloat64 `json:"mosh_m"`
// 	Microbio    sql.NullString  `json:"microbio"`
// 	S_cs137     sql.NullFloat64 `json:"s_cs137"`
// 	S_sr90      sql.NullFloat64 `json:"s_sr90"`
// 	Solon_d     sql.NullFloat64 `json:"solon_d"`
// 	Fact_isp    sql.NullString  `json:"fact_isp"`
// 	Eco_na      sql.NullFloat64 `json:"eco_na"`
// 	El_mn       sql.NullFloat64 `json:"el_mn"`
// 	El_co       sql.NullFloat64 `json:"el_co"`
// 	Pol_id      sql.NullInt64   `json:"pol_id"`
// 	Gosprog     sql.NullInt64   `json:"gosprog"`
// 	Rekisp      sql.NullInt64   `json:"rekisp"`
// 	Geom_json   string          `db:"geom_json" json:"geom_json"`
// 	Humus_class sql.NullString  `json:"humus_class"`
// 	Class_k     sql.NullString  `json:"class_k"`
// 	Class_p     sql.NullString  `json:"class_p"`
// 	Title       sql.NullString  `json:"title"`
// 	Title2      sql.NullString  `json:"Title2"`
// }
