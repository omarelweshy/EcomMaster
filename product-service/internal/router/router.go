package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omarelweshy/EcomMaster-product-service/internal/handler"
	"github.com/omarelweshy/EcomMaster-product-service/internal/repository"
	"github.com/omarelweshy/EcomMaster-product-service/internal/service"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	productRepository := &repository.ProductRepository{DB: db}
	productService := &service.ProductService{Repo: *productRepository}
	productHandler := &handler.ProductHandler{ProductService: productService}

	r := gin.Default()
	// r.Use(middleware.Logging())
	// r.Use(middleware.ErrorHandler())

	r.POST("/product/create", productHandler.CreateProduct)
	return r
}
