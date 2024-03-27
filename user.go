package gogis

type User struct {
	ID           int    `json:"-"`
	Username     string `json:"Username" binding:"required"`
	PasswordHash string `json:"-"`
	Email        string `json:"Email" binding:"required"`
}
