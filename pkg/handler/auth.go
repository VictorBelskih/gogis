package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/VictorBelskih/gogis"
	"github.com/VictorBelskih/gogis/pkg/render"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUpView(c *gin.Context) {
	users, err := h.services.Authorization.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	userrole, err := h.services.Authorization.GetRole()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select role"})
		return
	}
	data := gin.H{
		"users": users,
		"roles": userrole,
	}

	render.RenderTemplate(c, "signup", data)
}

func (h *Handler) signInView(c *gin.Context) {
	// Ваша логика обработки входа пользователя
	render.RenderTemplate(c, "signin", nil)
}
func (h *Handler) signUp(c *gin.Context) {
	var userdata gogis.User

	if c.Request.Method == "POST" {
		err := c.Request.ParseForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		username := c.PostForm("Username")
		email := c.PostForm("Email")
		password := c.PostForm("Password")
		roleStr := c.PostForm("Role")
		// Проверка наличия обязательных полей
		if username == "" || email == "" || password == "" || roleStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}

		// Преобразование строки role в целое число
		role, err := strconv.Atoi(roleStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role format"})
			return
		}

		// Валидация email и других полей, если необходимо

		userdata = gogis.User{
			Username:     username,
			Email:        email,
			PasswordHash: password,
			Role:         role, // Использование преобразованного значения role
		}

		userID, err := h.services.Authorization.CreateUser(userdata)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "userID": userID})
	}
}

func (h *Handler) signIn(c *gin.Context) {
	username := c.PostForm("Username")
	password := c.PostForm("Password")

	token, err := h.services.Authorization.AuthenticateUser(username, password)
	if err != nil {
		// Обработка ошибки аутентификации
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверное имя пользователя или пароль"})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   3600, // Время жизни cookie в секундах
		HttpOnly: true, // Флаг HttpOnly защищает cookie от доступа через JavaScript
		Secure:   true, // Флаг Secure гарантирует отправку cookie только по HTTPS
		Path:     "/",  // Cookie будет доступен для всего сайта
	})
	c.Redirect(http.StatusFound, "/")
	// Отправляем токен непосредственно клиенту в теле ответа
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Вход в систему успешен",
	// 	"token":   token,
	// })
}
func (h *Handler) signOut(c *gin.Context) {
	// Установка cookie с именем 'token' и сроком жизни в прошлом, чтобы удалить его
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0), // Дата в прошлом
		HttpOnly: true,            // Флаг HttpOnly защищает cookie от доступа через JavaScript
		Secure:   true,            // Флаг Secure гарантирует отправку cookie только по HTTPS
		Path:     "/",             // Cookie будет доступен для всего сайта
	})

	c.Redirect(http.StatusFound, "/auth/signin")
}
