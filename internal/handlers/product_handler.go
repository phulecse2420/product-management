package handlers

import (
	"errors"
	"net/http"
	"pm/internal/models"
	"pm/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{ service *services.ProductService }

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/products", h.Create)
	r.GET("/products", h.List)
	r.GET("/products/:id", h.GetByID)
	r.PUT("/products/:id", h.Update)
	r.DELETE("/products/:id", h.Delete)
}

func (h *ProductHandler) Create(context *gin.Context) {
	var inp models.CreateProductInput
	if err := context.ShouldBindJSON(&inp); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	p, err := h.service.Create(inp)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, p)
}

func (h *ProductHandler) List(c *gin.Context) {
	keyword := c.Query("keyword")
	list, err := h.service.List(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return
	}
	if list == nil {
		list = []models.Product{}
	}
	c.JSON(http.StatusOK, list)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	p, err := h.service.GetByID(id)
	if errors.Is(err, services.ErrProductNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	var inp models.UpdateProductInput
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	p, err := h.service.Update(id, inp)
	if errors.Is(err, services.ErrProductNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	if err := h.service.Delete(id); errors.Is(err, services.ErrProductNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
