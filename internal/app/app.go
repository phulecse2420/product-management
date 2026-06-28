package app

import (
	"pm/config"
	"pm/internal/handlers"
	"pm/internal/infrastructure"
	"pm/internal/repositories"
	"pm/internal/routes"
	"pm/internal/services"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) error {
	db, err := infrastructure.NewDB(cfg.DSN())
	if err != nil {
		return err
	}
	defer db.Close()

	productRepository := repositories.NewProductRepository(db)

	productService := services.NewProductService(productRepository)

	healthHandler := handlers.NewHealthHandler()
	productHandler := handlers.NewProductHandler(productService)

	r := gin.Default()
	routes.Register(r, productHandler, healthHandler)

	return r.Run(":" + cfg.Port)
}
