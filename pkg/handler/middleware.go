package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получение токена из cookie
		tokenCookie, err := c.Request.Cookie("token")
		if err != nil {
			// Перенаправление на страницу авторизации вместо отправки JSON ответа
			c.Redirect(http.StatusFound, "/auth/signin")
			return
		}
		tokenString := tokenCookie.Value

		// Парсинг и валидация токена
		user, err := h.services.Authorization.ParseJWTToken(tokenString)
		if err != nil {
			log.Println("Ошибка при парсинге токена:", err)
			// Перенаправление на страницу авторизации вместо отправки JSON ответа
			c.Redirect(http.StatusFound, "/auth/signin")
			return
		}

		// Установка информации о пользователе в контекст
		c.Set("userID", user.ID)
		c.Set("username", user.Username)
		c.Set("role", user.Role)

		// Передача управления следующему обработчику в цепочке
		c.Next()
	}
}
