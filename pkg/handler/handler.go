package handler

import (
	"os"

	"github.com/VictorBelskih/gogis/pkg/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func setupSessions(router *gin.Engine, secretKey string) {
	store := cookie.NewStore([]byte(secretKey))
	router.Use(sessions.Sessions("mysession", store))
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	secretKey := os.Getenv("SESSION_SECRET_KEY")
	if secretKey == "" {
		secretKey = "defaultSecretKey"
	}
	router := gin.New()
	setupSessions(router, secretKey)
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
	gis := router.Group("/")
	{
		gis.GET("/", h.gisPage)
		gis.GET("/gis/spr_cult", h.sprCult)
		gis.GET("/gis/spr_cult/addView", h.CultAddView)
		gis.POST("/gis/spr_cult/add", h.createCult)
		gis.POST("/gis/spr_cult/update/:id", h.createCult)
		gis.GET("/gis/spr_cult/del/:id", h.deleteCult)
	}
	//router.GET("/", h.homePage)

	return router
}
