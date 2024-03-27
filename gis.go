package gogis

type Field struct {
	// Id        int     `json:"id"`
	//Geom string `json:"geom"`
	Id_eu int `json:"id_eu"`
	// Station   string  `json:"station"`
	// Date      string  `json:"date"`
	// Sub_name  float64 `json:"sub_name"`
	// Reg_name  string  `json:"reg_name"`
	// Reg_id    string  `json:"reg_id"`
	// Reg_oktmo string  `json:"reg_oktmo"`
	// Sp_name   string  `json:"sp_name"`
	// Sp_id     string  `json:"sp_id"`
	// Date_dzz  string  `json:"date_dzz"`
	// Dzz_ka_s  string  `json:"dzz_ka_s"`
	Farm_name string  `json:"farm_name"`
	Tlu       float64 `json:"tlu"`
	// Mlrt_d    float64 `json:"mlrt_d"`
	// Mlrt_w    float64 `json:"mlrt_w"`
	Area_f float64 `json:"area_f"`
	// Crop      float64 `json:"crop"`
	// Crop_old  float64 `json:"crop_old"`
	// S_type    float64 `json:"s_type"`
	// S_ptype   float64 `json:"s_ptype"`
	// S_ms      float64 `json:"s_ms"`
	// S_ind     float64 `json:"s_ind"`
	// Organic   float64 `json:"organic"`
	// Ph_kcl    float64 `json:"ph_kcl"`
	// Ph_water  float64 `json:"ph_water"`
	// Gydr_c    float64 `json:"gydr_c"`
	// Eko       float64 `json:"eko"`
	// El_p      float64 `json:"el_p"`
	// El_k      float64 `json:"el_k"`
	// El_n      float64 `json:"el_n"`
	// El_ca     float64 `json:"el_ca"`
	// El_mg     float64 `json:"el_mg"`
	// El_zn     float64 `json:"el_zn"`
	// El_cu     float64 `json:"el_cu"`
	// El_mo     float64 `json:"el_mo"`
	// El_s      float64 `json:"el_s"`
	// El_b      float64 `json:"el_b"`
	// Il        float64 `json:"il"`
	// Glina     float64 `json:"glina"`
	// Torf      float64 `json:"torf"`
	// Pereyv    float64 `json:"pereyv"`
	// Plotn     float64 `json:"plotn"`
	// Agreg     float64 `json:"agreg"`
	// Field_vl  float64 `json:"field_vl"`
	// Nas_osn   float64 `json:"nas_osn"`
	// DGRD_ST   float64 `json:"dgrd_st"`
	// Yklon     float64 `json:"yklon"`
	// Metals    float64 `json:"metals"`
	// Met_pb    float64 `json:"met_pb"`
	// Met_cd    float64 `json:"met_cd"`
	// Met_hg    float64 `json:"met_hg"`
	// Met_as    float64 `json:"met_as"`
	// Plt_tox   float64 `json:"plt_tox"`
	// Plt_oil   float64 `json:"plt_oil"`
	// Cs137     float64 `json:"cs137"`
	// Sr90      float64 `json:"sr90"`
	// Gamma     float64 `json:"gamma"`
	// Salinity  float64 `json:"salinity"`
	// Solon_Na  float64 `json:"solon_na"`
	// Solon_pr  float64 `json:"solon_pr"`
	// Wind_er   float64 `json:"wind_er"`
	// Water_er  float64 `json:"water_er"`
	// Zakust    float64 `json:"zakust"`
	// Zales     float64 `json:"zales"`
	// Zarast    float64 `json:"zarast"`
	// Zastroen  float64 `json:"zastroen"`
	// Narushen  float64 `json:"narushen"`
	// Prochee   float64 `json:"prochee"`
	// Mosh_g    float64 `json:"mosh_g"`
	// Mosh_m    float64 `json:"mosh_m"`
	// Microbio  string  `json:"microbio"`
	// S_cs137   float64 `json:"s_cs137"`
	// S_sr90    float64 `json:"s_sr90"`
	// Solon_d   float64 `json:"solon_d"`
	// Fact_isp  string  `json:"fact_isp"`
	// Eco_na    float64 `json:"eco_na"`
	// El_mn     float64 `json:"el_mn"`
	// El_co     float64 `json:"el_co"`
	// Pol_id    int     `json:"pol_id"`
	// Gosprog   int     `json:"gosprog"`
	// Rekisp    int     `json:"rekisp"`
	Geom_json string `db:"geom_json" json:"geom_json"`
}