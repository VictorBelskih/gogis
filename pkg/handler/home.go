package handler

import (
	"log"
	"net/http"

	"github.com/VictorBelskih/gogis/pkg/render"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) homePage(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("token")
	if token == nil {
		// Обработка отсутствия токена, например, перенаправление на страницу входа
		c.Redirect(http.StatusFound, "/auth/signin")
		return
	}

	authUser, err := h.services.Authorization.ParseJWTToken(token.(string))
	if err != nil {
		// Обработка ошибки парсинга токена, например, логирование или отображение сообщения пользователю
		log.Println("Ошибка при парсинге токена:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fields, err := h.services.Gis.GetField()
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	data := gin.H{
		"token":    token,
		"id":       authUser.ID,
		"username": authUser.Username,
		"fields":   fields,
	}

	render.RenderTemplate(c, "index", data)
}
