package handler

import (
	"github.com/VictorBelskih/gogis/pkg/render"
	"github.com/gin-gonic/gin"
)

func (h *Handler) homePage(c *gin.Context) {
	// Логика для отображения главной страницы
	render.RenderTemplate(c, "index", nil)
}
