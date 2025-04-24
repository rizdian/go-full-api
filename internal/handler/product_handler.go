package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rizdian/go-full-api/internal/model"
	"github.com/rizdian/go-full-api/internal/repository"
	"github.com/rizdian/go-full-api/internal/service"
	"github.com/rizdian/go-full-api/internal/utils"
	"gorm.io/gorm"
	"net/http"
)

func RegisterProductRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)

	rg.POST("/products", func(c *gin.Context) {
		var product model.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			utils.BadRequestResponse(c, err.Error())
			return
		}

		if err := productService.Create(&product); err != nil {
			utils.InternalServerErrorResponse(c, err.Error())
			return
		}

		utils.SuccessResponse(c, http.StatusCreated, "Product created successfully", product)
	})

	rg.GET("/products", func(c *gin.Context) {
		products, err := productService.GetAll()
		if err != nil {
			utils.InternalServerErrorResponse(c, err.Error())
			return
		}

		utils.SuccessResponse(c, http.StatusOK, "Products fetched successfully", products)
	})
}
