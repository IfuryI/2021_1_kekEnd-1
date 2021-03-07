package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/auth"
)

func RegisterHttpEndpoints(router *gin.Engine, authUC auth.UseCase) {
	handler := NewHandler(authUC)

	router.POST("/users", handler.CreateUser)
	router.GET("/users/:id", handler.GetUser)
	router.PUT("/users/:id", handler.UpdateUser)
}
