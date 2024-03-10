package handler

import (
	"github.com/VictorBelskih/gogis/pkg/render"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUpView(c *gin.Context) {
	// Ваша логика обработки регистрации пользователя
	render.RenderTemplate(c, "signup", nil)
}

func (h *Handler) signInView(c *gin.Context) {
	// Ваша логика обработки входа пользователя
	render.RenderTemplate(c, "signin", nil)
}
func (h *Handler) signUp(c *gin.Context) {

}

func (h *Handler) signIn(c *gin.Context) {

}
