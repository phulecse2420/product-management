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

	repo := repositories.NewProductRepository(db)

	svc := services.NewProductService(repo)

	handler := handlers.NewProductHandler(svc)

	r := gin.Default()
	routes.Register(r, handler)
	return r.Run(":" + cfg.Port)
}
