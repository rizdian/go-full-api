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

func RegisterOrderRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	orderRepo := repository.NewOrderRepository(db)
	productRepo := repository.NewProductRepository(db)
	orderService := service.NewOrderService(orderRepo, productRepo)

	rg.POST("/orders", func(c *gin.Context) {
		var order model.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			utils.BadRequestResponse(c, err.Error())
			return
		}

		// Menyimpan order
		if err := orderService.Create(&order); err != nil {
			utils.InternalServerErrorResponse(c, err.Error())
			return
		}

		utils.SuccessResponse(c, http.StatusCreated, "Order created successfully", order)
	})

	rg.GET("/orders", func(c *gin.Context) {
		orders, err := orderService.GetAll()
		if err != nil {
			utils.InternalServerErrorResponse(c, err.Error())
			return
		}

		utils.SuccessResponse(c, http.StatusOK, "Orders fetched successfully", orders)
	})
}
