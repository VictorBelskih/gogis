package repository

import (
	"github.com/VictorBelskih/gogis"
	"github.com/jmoiron/sqlx"
)

// связь с бд
type AuthPostgres struct {
	db *sqlx.DB
}

// создание экземпляра репозитория
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// добавление пользователя
func (r *AuthPostgres) CreateUser(user gogis.User) (int, error) {
	query := "INSERT INTO users (username, password, email, role) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := r.db.QueryRow(query, user.Username, user.PasswordHash, user.Email, user.Role).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// получение пользователя для отладки
func (r *AuthPostgres) GetUsers() ([]gogis.User, error) {
	query := "SELECT id, username FROM users"
	users := []gogis.User{}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user gogis.User
		err = rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *AuthPostgres) GetRole() ([]gogis.Role, error) {
	query := "SELECT * FROM role"
	roles := []gogis.Role{}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var role gogis.Role
		err = rows.Scan(&role.Id, &role.Role)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

// получение пользователя
func (r *AuthPostgres) GetUserByUsername(username string) (gogis.User, error) {
	query := `SELECT id, username, password, email, role FROM users WHERE username = $1`
	var user gogis.User

	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Email, &user.Role)
	if err != nil {
		return gogis.User{}, err
	}

	return user, nil
}
