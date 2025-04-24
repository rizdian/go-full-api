package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rizdian/go-full-api/internal/model"
	"github.com/rizdian/go-full-api/internal/utils"
	"gorm.io/gorm"
	"net/http"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	rg.POST("/login", func(c *gin.Context) {
		var req struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required"`
		}

		// Bind incoming request JSON to struct
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.BadRequestResponse(c, err.Error())
			return
		}

		// Fetch user from database by email
		var user model.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			utils.UnauthorizedResponse(c, "Invalid credentials")
			return
		}

		// Compare plain text passwords (later we will implement hashed passwords)
		if user.Password != req.Password {
			utils.UnauthorizedResponse(c, "Invalid credentials")
			return
		}

		// Generate JWT token
		token, err := utils.GenerateToken(user.ID)
		if err != nil {
			utils.InternalServerErrorResponse(c, "Failed to generate token")
			return
		}

		// Respond with the token
		utils.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{
			"token": token,
		})
	})
}
