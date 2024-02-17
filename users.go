package gogis

type Users struct {
	id       int    `json:"-"`
	username string `json:"username"`
	password string `json:"password"`
}
