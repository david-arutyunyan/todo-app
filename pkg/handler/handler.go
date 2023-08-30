package handler

import (
	"github.com/gin-gonic/gin"
	"github/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine { // Инициализатор всех энд-поинтов
	router := gin.New()

	usersSegService := router.Group("/users-segmentation")
	{
		segments := usersSegService.Group("/segment")
		{
			segments.POST("/", h.createSegment)
			segments.DELETE("/", h.deleteSegment)
		}

		users := usersSegService.Group("/user")
		{
			users.POST("/", h.createUser)
			users.DELETE("/:id", h.deleteUser)
		}

		usersSegments := usersSegService.Group("/users-segments")
		{
			usersSegments.POST("/", h.updateUserSegments)
			usersSegments.GET("/", h.getUserSegments)
		}

	}

	return router
}
