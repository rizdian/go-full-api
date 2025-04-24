package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SuccessResponse digunakan untuk mengirimkan respons berhasil
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

// ErrorResponse digunakan untuk mengirimkan respons error
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  "error",
		"message": message,
	})
}

// NotFoundResponse digunakan untuk mengirimkan respons jika resource tidak ditemukan
func NotFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  "error",
		"message": message,
	})
}

// BadRequestResponse digunakan untuk mengirimkan respons jika ada permintaan yang tidak valid
func BadRequestResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "error",
		"message": message,
	})
}

// InternalServerErrorResponse digunakan untuk mengirimkan respons jika terjadi kesalahan server
func InternalServerErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  "error",
		"message": message,
	})
}

// UnauthorizedResponse digunakan untuk mengirimkan respons jika kredensial tidak valid
func UnauthorizedResponse(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  "error",
		"message": message,
	})
}
