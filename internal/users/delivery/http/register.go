package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/logger"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/middleware"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/sessions"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/users"
)

func RegisterHttpEndpoints(router *gin.Engine, usersUC users.UseCase, sessions sessions.Delivery,
	authMiddleware middleware.Auth, Log *logger.Logger) {
	handler := NewHandler(usersUC, sessions, Log)

	router.POST("/users", handler.CreateUser)
	router.POST("/users/avatar", authMiddleware.CheckAuth(), handler.UploadAvatar)
	router.GET("/users", authMiddleware.CheckAuth(), handler.GetUser)
	router.PUT("/users", authMiddleware.CheckAuth(), handler.UpdateUser)
	router.DELETE("/sessions", authMiddleware.CheckAuth(), handler.Logout)
	router.POST("/sessions", handler.Login)

	router.POST("/subscribe/:user_id", authMiddleware.CheckAuth(), handler.Subscribe)
	router.DELETE("/subscribe/:user_id", authMiddleware.CheckAuth(), handler.Unsubscribe)
	router.GET("/subscribers/:user_id", handler.GetSubscribers)
	router.GET("/subscriptions/:user_id", handler.GetSubscriptions)
}
