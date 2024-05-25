package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/VictorBelskih/gogis"
	"github.com/VictorBelskih/gogis/pkg/render"
	"github.com/gin-gonic/gin"
)

func (h *Handler) gisPage(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		// userID не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить userID"})
		return
	}
	// Приведение типа к int
	intUserID, ok := userID.(int)
	if !ok {
		// userID не может быть приведен к типу int, обработка ошибки
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID должен быть целым числом"})
		return
	}
	username, exists := c.Get("username")
	if !exists {
		// username не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить username"})
		return
	}
	role, exists := c.Get("role")
	if !exists {
		// username не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить role"})
		return
	}
	intUserRole, ok := role.(int)
	if !ok {
		// userID не может быть приведен к типу int, обработка ошибки
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID должен быть целым числом"})
		return
	}
	userField, err := h.services.Gis.GetFieldByUser(intUserID, intUserRole)
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
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
	FieldData, err := h.services.Gis.GetFieldData(intUserID, intUserRole)
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	totalAreaByFieldType, err := h.services.Gis.CalculateTotalAreaByFieldType(intUserID, intUserRole)
	if err != nil {
		log.Println("Error Calculation Area", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	totalArea, err := h.services.Gis.TotalArea(intUserID, intUserRole)
	if err != nil {
		log.Println("Error Calculation Area", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	avgOrganic, err := h.services.Gis.CalculateAverageHumusByClass(intUserID, intUserRole)
	if err != nil {
		log.Println("Error Calculation AVG Organic", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	AvgK, err := h.services.Gis.AvgPotassiumByClass(intUserID, intUserRole)
	if err != nil {
		log.Println("Error Calculation AVG k", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	AvgP, err := h.services.Gis.AvgPhosphorByClass(intUserID, intUserRole)
	if err != nil {
		log.Println("Error Calculation AVG p", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	avgRad, err := h.services.Gis.CalculateRadionuclideSummary(intUserID, intUserRole)
	data := gin.H{
		"userID":               userID,
		"username":             username,
		"role":                 role,
		"fields":               fields,
		"FieldData":            FieldData,
		"totalAreaByFieldType": totalAreaByFieldType,
		"totalArea":            totalArea,
		"AvgOrganic":           avgOrganic,
		"avgRad":               avgRad,
		"AvgK":                 AvgK,
		"AvgP":                 AvgP,
		"userField":            userField,
	}

	render.RenderTemplate(c, "index", data)
}
func (h *Handler) sprCult(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		// userID не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить userID"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		// username не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить username"})
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
		"userID":   userID,
		"username": username,
		"cult":     cultData,
	}
	render.RenderTemplate(c, "spr_cult", data)
}

func (h *Handler) CultAddView(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		// userID не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить userID"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		// username не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить username"})
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
		"userID":   userID,
		"username": username,
		"cult":     cultData,
	}
	render.RenderTemplate(c, "cult_add", data)
}

func (h *Handler) CultUpdateView(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		// userID не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить userID"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		// username не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить username"})
		return
	}
	idStr := c.Param("id")

	// Преобразование id из строки в int
	id, err := strconv.Atoi(idStr)
	cultData, err := h.services.Gis.GetCultByID(id)
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	data := gin.H{
		"userID":   userID,
		"username": username,
		"cult":     cultData,
	}
	render.RenderTemplate(c, "cult_update", data)
}

func (h *Handler) updateCult(c *gin.Context) {
	var cultData gogis.Cult

	if c.Request.Method == "POST" {
		err := c.Request.ParseForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		oldIdStr := c.PostForm("old_cult_id") // Получаем старый ID из формы
		newIdStr := c.PostForm("cult_id")     // Получаем новый ID из формы
		title := c.PostForm("cult_name")

		// Проверка на отсутствие необходимых полей
		if oldIdStr == "" || newIdStr == "" || title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Отсутствуют необходимые поля"})
			return
		}

		// Конвертация oldId и newId из строки в int
		oldId, err := strconv.Atoi(oldIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат старого ID"})
			return
		}

		newId, err := strconv.Atoi(newIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат нового ID"})
			return
		}

		// Валидация полей, если необходимо

		cultData = gogis.Cult{
			OldId: oldId, // Устанавливаем старый ID
			Id:    newId, // Устанавливаем новый ID
			Title: title,
		}

		err = h.services.Gis.UpdateCult(cultData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить данные культуры"})
			return
		}

		// Перенаправление на страницу '/gis/spr_cult'
		c.Redirect(http.StatusFound, "/gis/spr_cult")
	}
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
			OldId: id,
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

func (h *Handler) sprFarm(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		// userID не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить userID"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		// username не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить username"})
		return
	}

	farmData, err := h.services.Gis.GetFarm()
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	data := gin.H{
		"userID":   userID,
		"username": username,
		"farm":     farmData,
	}
	render.RenderTemplate(c, "spr_farm", data)
}

func (h *Handler) FarmAddView(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		// userID не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить userID"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		// username не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить username"})
		return
	}
	farmData, err := h.services.Gis.GetFarm()
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	districtData, err := h.services.Gis.GetDistrict()
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	userData, err := h.services.Authorization.GetUsers()
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	data := gin.H{
		"userID":   userID,
		"username": username,
		"district": districtData,
		"farm":     farmData,
		"user":     userData,
	}
	render.RenderTemplate(c, "farm_add", data)
}

func (h *Handler) CreateFarm(c *gin.Context) {
	var farmData gogis.Farm

	if c.Request.Method == "POST" {
		err := c.Request.ParseForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		idStr := c.PostForm("farm_id")
		Name := c.PostForm("farm_name")
		DistrictStr := c.PostForm("district")
		UserStr := c.PostForm("user")
		// Check for missing required fields
		if idStr == "" || Name == "" || DistrictStr == "" || UserStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}
		User, err := strconv.Atoi(UserStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		District, err := strconv.Atoi(DistrictStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		// Convert id from string to int
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		// Validate fields if necessary

		farmData = gogis.Farm{
			OldId:    id,
			Id:       id,
			Name:     Name,
			District: District,
			Id_user:  User,
		}

		err = h.services.Gis.CreateFarm(farmData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cult"})
			return
		}

		// Redirect to the '/gis/spr_cult' page
		c.Redirect(http.StatusFound, "/gis/spr_farm")
	}
}
func (h *Handler) deleteFarm(c *gin.Context) {
	// Извлечение значения id из параметра пути
	idStr := c.Param("id")

	// Преобразование id из строки в int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Вызов функции удаления с переданным id
	err = h.services.Gis.DeleteFarm(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cult"})
		return
	}
	// Redirect to the '/gis/spr_cult' page
	c.Redirect(http.StatusFound, "/gis/spr_farm")
}
func (h *Handler) FarmUpdateView(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		// userID не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить userID"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		// username не найден в контексте, обработка ошибки
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить username"})
		return
	}
	idStr := c.Param("id")

	// Преобразование id из строки в int
	id, err := strconv.Atoi(idStr)
	farmData, err := h.services.Gis.GetFarmByID(id)
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	districtData, err := h.services.Gis.GetDistrict()
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	userData, err := h.services.Authorization.GetUsers()
	if err != nil {
		// Handle error fetching fields, for example, log or display a message to the user
		log.Println("Error fetching fields:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	data := gin.H{
		"userID":   userID,
		"username": username,
		"district": districtData,
		"farm":     farmData,
		"user":     userData,
	}
	render.RenderTemplate(c, "farm_update", data)
}
func (h *Handler) updateFarm(c *gin.Context) {
	var farmData gogis.Farm

	if c.Request.Method == "POST" {
		err := c.Request.ParseForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		oldIdStr := c.PostForm("old_id") // Объявление переменной oldIdStr
		idStr := c.PostForm("farm_id")
		Name := c.PostForm("farm_name")
		DistrictStr := c.PostForm("district")
		UserStr := c.PostForm("user")

		// Проверка на отсутствие необходимых полей
		if oldIdStr == "" || idStr == "" || Name == "" || DistrictStr == "" || UserStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Отсутствуют необходимые поля"})
			return
		}
		User, err := strconv.Atoi(UserStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		District, err := strconv.Atoi(DistrictStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		// Конвертация oldId и newId из строки в int
		oldId, err := strconv.Atoi(oldIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат старого ID"})
			return
		}

		newId, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат нового ID"})
			return
		}

		// Валидация полей, если необходимо

		farmData = gogis.Farm{
			OldId:    oldId,
			Id:       newId,
			Name:     Name,
			District: District,
			Id_user:  User,
		}

		err = h.services.Gis.UpdateFarm(farmData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить данные культуры"})
			return
		}

		// Перенаправление на страницу '/gis/spr_cult'
		c.Redirect(http.StatusFound, "/gis/spr_farm")
	}
}
