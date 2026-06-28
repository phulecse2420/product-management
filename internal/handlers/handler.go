package handlers

import "github.com/gin-gonic/gin"

type Registerable interface {
	RegisterRoutes(r *gin.Engine)
}
