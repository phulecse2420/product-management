package routes

import (
	"pm/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, handlers ...handlers.Registerable) {
	for _, h := range handlers {
		h.RegisterRoutes(r)
	}
}
