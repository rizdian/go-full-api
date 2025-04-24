package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rizdian/go-full-api/internal/model"
	"github.com/rizdian/go-full-api/internal/repository"
	"github.com/rizdian/go-full-api/internal/service"
	"github.com/rizdian/go-full-api/internal/utils" // Impor utils sesuai dengan path yang benar
	"gorm.io/gorm"
	"net/http"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	rg.POST("/users", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			utils.BadRequestResponse(c, err.Error()) // Menggunakan utils untuk mengirimkan response error
			return
		}

		if err := userService.Create(&user); err != nil {
			utils.InternalServerErrorResponse(c, err.Error()) // Menggunakan utils untuk mengirimkan response error
			return
		}

		utils.SuccessResponse(c, http.StatusCreated, "User created successfully", user) // Response sukses
	})

	rg.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, err := userService.GetByID(id)
		if err != nil {
			utils.NotFoundResponse(c, "User not found") // Menggunakan utils untuk mengirimkan response error 404
			return
		}

		utils.SuccessResponse(c, http.StatusOK, "User fetched successfully", user) // Response sukses
	})
}
