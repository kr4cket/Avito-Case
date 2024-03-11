package handlers

import (
	"avitoCase/pkg/services"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "avitoCase/docs"
)

type Handler struct {
	services *services.Service
}

func New(service *services.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("/api")
	{
		cost := api.Group("/service")
		{
			cost.GET("/cost/:id", h.getCost)
		}
	}

	return router
}
