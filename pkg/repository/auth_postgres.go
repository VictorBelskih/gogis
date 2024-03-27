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
	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := r.db.QueryRow(query, user.Username, user.PasswordHash, user.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// получение пользователя для отладки
func (r *AuthPostgres) GetUsers() ([]gogis.User, error) {
	query := "SELECT id, username, password email FROM users"
	users := []gogis.User{}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user gogis.User
		err = rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// получение пользователя
func (r *AuthPostgres) GetUserByUsername(username string) (gogis.User, error) {
	query := "SELECT id, username, password, email FROM users WHERE username = $1"
	var user gogis.User

	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Email)
	if err != nil {
		return gogis.User{}, err
	}

	return user, nil
}
