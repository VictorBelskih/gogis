package handler

import (
	"net/http"

	"github.com/VictorBelskih/gogis"
	"github.com/VictorBelskih/gogis/pkg/render"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUpView(c *gin.Context) {
	users, err := h.services.Authorization.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	data := gin.H{
		"users": users,
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

		// Проверка наличия обязательных полей
		if username == "" || email == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}

		// Валидация email и других полей, если необходимо

		userdata = gogis.User{
			Username:     username,
			Email:        email,
			PasswordHash: password,
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
	session := sessions.Default(c)
	session.Set("token", token)
	session.Save()
	c.Redirect(http.StatusFound, "/")
	//c.JSON(http.StatusOK, gin.H{"message": "Вход в систему успешен", "token": token})
}
func (h *Handler) signOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("token") // Удаление токена из сессии
	session.Save()

	c.Redirect(http.StatusFound, "/")
}
