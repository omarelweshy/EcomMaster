package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/omarelweshy/EcomMaster-product-service/internal/form"
	"github.com/omarelweshy/EcomMaster-product-service/internal/service"
	"github.com/omarelweshy/EcomMaster-product-service/internal/util"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var form form.CreateProductForm

	if err := c.ShouldBindJSON(&form); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := util.FormatValidationError(validationErrors)
			util.RespondWithError(c, http.StatusBadRequest, "Validation failed", formattedErrors)
			return
		}
		util.RespondWithError(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	err := h.ProductService.CreateProduct(form.Name, form.Description, form.Price, form.Stock) // We will update category later on
	if err != nil {
		util.RespondWithError(c, http.StatusBadRequest, "Unable to create product", nil)
		return
	}
	util.RespondWithSuccess(c, "product created", nil)
	return
}
