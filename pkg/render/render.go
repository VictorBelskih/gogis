package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderTemplate(c *gin.Context, tmpl string, data interface{}) {
	c.HTML(http.StatusOK, tmpl+".html", data)
}

// package render

// import (
// 	"html/template"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // Функция для рендеринга HTML страницы с подключением header
// func RenderTemplate(c *gin.Context, tmpl string, data interface{}) {
// 	// Загрузка всех шаблонов
// 	var tmplList = []string{
// 		"templates/base.html",
// 		"templates/header.html",
// 		"templates/" + tmpl + ".html",
// 	}
// 	var html = template.Must(template.ParseFiles(tmplList...))
// 	c.SetHTMLTemplate(html)

// 	// Рендеринг страницы с подключенным header
// 	c.HTML(http.StatusOK, tmpl+".html", data)
// }
