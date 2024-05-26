package handler

import (
	"os"

	"github.com/VictorBelskih/gogis/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

// func setupSessions(router *gin.Engine, secretKey string) {
// 	store := cookie.NewStore([]byte(secretKey))
// 	router.Use(sessions.Sessions("mysession", store))
// }

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	secretKey := os.Getenv("SESSION_SECRET_KEY")
	if secretKey == "" {
		secretKey = "defaultSecretKey"
	}
	router := gin.New()
	// setupSessions(router, secretKey)
	// router.Use(middleware.TokenAuthMiddleware())
	router.LoadHTMLGlob("templates/*.html")

	router.Static("/static", "./static")

	auth := router.Group("/auth")
	{
		auth.GET("/signin", h.signInView)
		auth.GET("/signup", h.signUpView)
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/signout", h.signOut)
	}
	gis := router.Group("/", h.TokenAuthMiddleware())
	{
		gis.GET("/", h.gisPage)
		gis.GET("/report", h.Report)
		gis.GET("/gis/spr_cult", h.sprCult)
		gis.GET("/gis/spr_cult/addView", h.CultAddView)
		gis.POST("/gis/spr_cult/add", h.createCult)
		gis.GET("/gis/spr_cult/updateView/:id", h.CultUpdateView)
		gis.POST("/gis/spr_cult/update", h.updateCult)
		gis.GET("/gis/spr_cult/del/:id", h.deleteCult)
		gis.GET("/gis/spr_farm", h.sprFarm)
		gis.GET("/gis/spr_farm/addView", h.FarmAddView)
		gis.GET("/gis/spr_farm/updateView/:id", h.FarmUpdateView)
		gis.POST("/gis/spr_farm/add", h.CreateFarm)
		gis.POST("/gis/spr_farm/update", h.updateFarm)
		gis.GET("/gis/spr_farm/del/:id", h.deleteFarm)
	}
	//router.GET("/", h.homePage)

	return router
}
