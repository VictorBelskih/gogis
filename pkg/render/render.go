package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderTemplate(c *gin.Context, tmpl string, data interface{}) {
	c.HTML(http.StatusOK, tmpl+".html", data)
}
