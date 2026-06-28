package routes

import (
	"pm/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, h *handlers.ProductHandler) {
	r.POST("/products", h.Create)
	r.GET("/products", h.List)
	r.GET("/products/:id", h.GetByID)
	r.PUT("/products/:id", h.Update)
	r.DELETE("/products/:id", h.Delete)
}
