package service

import (
	"errors"
	"os"
	"time"

	"github.com/VictorBelskih/gogis"
	"github.com/VictorBelskih/gogis/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// передача интерфейсов
type AuthService struct {
	repo repository.Authorization
}

// создание интерфейса интерфейсов
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// хеширование пароля
func (s *AuthService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// обработка регистрации
func (s *AuthService) CreateUser(user gogis.User) (int, error) {
	hashedPassword, err := s.hashPassword(user.PasswordHash)
	if err != nil {
		return 0, err
	}

	user.PasswordHash = hashedPassword
	return s.repo.CreateUser(user)
}

// получение пользователя отладка
func (s *AuthService) GetUsers() ([]gogis.User, error) {
	return s.repo.GetUsers()
}
func (s *AuthService) GetRole() ([]gogis.Role, error) {
	return s.repo.GetRole()
}

// сравнение паролей
func (s *AuthService) ComparePasswords(inputPassword, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) GenerateJWTToken(user gogis.User) (string, error) {
	// Создание нового JWT токена
	token := jwt.New(jwt.SigningMethodHS256)
	secretKey := os.Getenv("JWT_SECRET")
	// Установка данных пользователя в payload токена
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["role"] = user.Role                            // Добавление роли пользователя
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix() // Установка срока действия токена на 12 часов

	// Подпись токена с секретным ключом
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) ParseJWTToken(tokenString string) (gogis.User, error) {
	// Проверка секретного ключа
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return gogis.User{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Извлечение данных пользователя из токена
		userIDFloat64, ok := claims["id"].(float64)
		if !ok {
			return gogis.User{}, errors.New("Неверный формат ID в токене")
		}
		userID := int(userIDFloat64)

		username := claims["username"].(string)
		email := claims["email"].(string)

		roleFloat64, ok := claims["role"].(float64) // Извлечение роли пользователя как float64
		if !ok {
			return gogis.User{}, errors.New("Неверный формат роли в токене")
		}
		role := int(roleFloat64) // Преобразование роли в int

		user := gogis.User{
			ID:       userID,
			Username: username,
			Email:    email,
			Role:     role, // Установка роли пользователя
		}

		return user, nil
	} else {
		return gogis.User{}, errors.New("Неверный токен")
	}
}

// авторизация пользователя
func (s *AuthService) AuthenticateUser(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = s.ComparePasswords(password, user.PasswordHash)
	if err != nil {
		return "", err
	}

	// Генерация JWT токена
	token, err := s.GenerateJWTToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
