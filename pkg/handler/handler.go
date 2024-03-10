package handler

import (
	"github.com/VictorBelskih/gogis/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("templates/*.html")

	router.Static("/static", "./static")

	auth := router.Group("/auth")
	{
		auth.GET("/signin", h.signInView)
		auth.GET("/signup", h.signUpView)
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	router.GET("/", h.homePage)

	return router
}
