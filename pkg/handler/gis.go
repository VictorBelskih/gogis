package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/VictorBelskih/gogis"
	"github.com/VictorBelskih/gogis/pkg/render"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) gisPage(c *gin.Context) {
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
		c.Redirect(http.StatusFound, "/auth/signin")
		return
	}
	fields, err := h.services.Gis.GetField()
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	FieldData, err := h.services.Gis.GetFieldData()
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	totalAreaByFieldType, err := h.services.Gis.CalculateTotalAreaByFieldType()
	if err != nil {
		log.Println("Error Calculation Area", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	totalArea, err := h.services.Gis.TotalArea()
	if err != nil {
		log.Println("Error Calculation Area", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	avgOrganic, err := h.services.Gis.CalculateAverageHumusByClass()
	if err != nil {
		log.Println("Error Calculation AVG Organic", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	AvgK, err := h.services.Gis.AvgPotassiumByClass()
	if err != nil {
		log.Println("Error Calculation AVG k", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	AvgP, err := h.services.Gis.AvgPhosphorByClass()
	if err != nil {
		log.Println("Error Calculation AVG p", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	avgRad, err := h.services.Gis.CalculateRadionuclideSummary()
	data := gin.H{
		"token":                token,
		"id":                   authUser.ID,
		"username":             authUser.Username,
		"fields":               fields,
		"FieldData":            FieldData,
		"totalAreaByFieldType": totalAreaByFieldType,
		"totalArea":            totalArea,
		"AvgOrganic":           avgOrganic,
		"avgRad":               avgRad,
		"AvgK":                 AvgK,
		"AvgP":                 AvgP,
	}

	render.RenderTemplate(c, "index", data)
}
func (h *Handler) sprCult(c *gin.Context) {
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
		c.Redirect(http.StatusFound, "/auth/signin")
		return
	}

	cultData, err := h.services.Gis.GetCult()
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
		"cult":     cultData,
	}
	render.RenderTemplate(c, "spr_cult", data)
}

func (h *Handler) CultAddView(c *gin.Context) {
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
		c.Redirect(http.StatusFound, "/auth/signin")
		return
	}

	cultData, err := h.services.Gis.GetCult()
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
		"cult":     cultData,
	}
	render.RenderTemplate(c, "cult_add", data)
}
func (h *Handler) sprCultAdd(c *gin.Context) {
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
		c.Redirect(http.StatusFound, "/auth/signin")
		return
	}

	data := gin.H{
		"token":    token,
		"id":       authUser.ID,
		"username": authUser.Username,
	}
	render.RenderTemplate(c, "spr_cult", data)
}

func (h *Handler) createCult(c *gin.Context) {
	var cultData gogis.Cult

	if c.Request.Method == "POST" {
		err := c.Request.ParseForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		idStr := c.PostForm("cult_id")
		title := c.PostForm("cult_name")

		// Check for missing required fields
		if idStr == "" || title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}

		// Convert id from string to int
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		// Validate fields if necessary

		cultData = gogis.Cult{
			Id:    id,
			Title: title,
		}

		err = h.services.Gis.CreateCult(cultData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cult"})
			return
		}

		// Redirect to the '/gis/spr_cult' page
		c.Redirect(http.StatusFound, "/gis/spr_cult")
	}
}

func (h *Handler) deleteCult(c *gin.Context) {
	// Извлечение значения id из параметра пути
	idStr := c.Param("id")

	// Преобразование id из строки в int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Вызов функции удаления с переданным id
	err = h.services.Gis.DeleteCult(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cult"})
		return
	}
	// Redirect to the '/gis/spr_cult' page
	c.Redirect(http.StatusFound, "/gis/spr_cult")
}
