package controller

import (
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		ProductUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, error := p.ProductUsecase.GetProducts()

	if error != nil {
		ctx.JSON(http.StatusInternalServerError, error)
	}

	ctx.JSON(http.StatusOK, products)
}
