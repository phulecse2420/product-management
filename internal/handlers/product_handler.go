package handlers

import (
	"errors"
	"net/http"
	"pm/internal/models"
	"pm/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{ svc *services.ProductService }

func NewProductHandler(svc *services.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var inp models.CreateProductInput
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	p, err := h.svc.Create(inp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func (h *ProductHandler) List(c *gin.Context) {
	keyword := c.Query("keyword")
	list, err := h.svc.List(keyword)
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
	p, err := h.svc.GetByID(id)
	if errors.Is(err, services.ErrNotFound) {
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
	p, err := h.svc.Update(id, inp)
	if errors.Is(err, services.ErrNotFound) {
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
	if err := h.svc.Delete(id); errors.Is(err, services.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
