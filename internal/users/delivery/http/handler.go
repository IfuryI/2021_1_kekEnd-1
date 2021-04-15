package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/logger"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/models"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/sessions"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/users"
	"github.com/go-park-mail-ru/2021_1_kekEnd/pkg/const"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
)

type Handler struct {
	useCase  users.UseCase
	sessions sessions.Delivery
	Log *logger.Logger
}

func NewHandler(useCase users.UseCase, sessions sessions.Delivery, Log *logger.Logger) *Handler {
	return &Handler{
		useCase:  useCase,
		sessions: sessions,
	}
}

type signupData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	signupData := new(signupData)

	err := ctx.BindJSON(signupData)
	if err != nil {
		msg := "Failed to bind request data " + err.Error()
		h.Log.LogWarning(ctx, "users", "CreateUser", msg)
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	if signupData.Username == "" || signupData.Email == "" || signupData.Password == "" {
		err := fmt.Errorf("%s", "invalid value in user data")
		h.Log.LogWarning(ctx, "users", "CreateUser", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	user := &models.User{
		Username:      signupData.Username,
		Email:         signupData.Email,
		Password:      signupData.Password,
		Avatar:        _const.DefaultAvatarPath,
		MoviesWatched: new(uint),
		ReviewsNumber: new(uint),
	}

	err = h.useCase.CreateUser(user)
	if err != nil {
		h.Log.LogError(ctx, "users", "CreateUser", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	userSessionID, err := h.sessions.Create(signupData.Username, _const.CookieExpires)
	if err != nil {
		h.Log.LogError(ctx, "users", "CreateUser", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	ctx.SetCookie(
		"session_id",
		userSessionID,
		int(_const.CookieExpires),
		"/",
		_const.Host,
		false,
		true,
	)

	ctx.Status(http.StatusCreated) // 201
}

func (h *Handler) Logout(ctx *gin.Context) {
	cookie, err := ctx.Cookie("session_id")
	if err != nil {
		h.Log.LogWarning(ctx, "users", "Logout", err.Error())
		ctx.AbortWithStatus(http.StatusUnauthorized) // 401
		return
	}

	err = h.sessions.Delete(cookie)
	if err != nil {
		h.Log.LogError(ctx, "users", "Logout", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	ctx.SetCookie("session_id", "Delete cookie", -1, "/", _const.Host, false, true)

	ctx.Status(http.StatusOK) // 200
}

type loginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Login(ctx *gin.Context) {
	loginData := new(loginData)

	err := ctx.BindJSON(loginData)
	if err != nil {
		msg := "Failed to bind request data " + err.Error()
		h.Log.LogWarning(ctx, "users", "Login", msg)
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	loginStatus := h.useCase.Login(loginData.Username, loginData.Password)
	if !loginStatus {
		err := fmt.Errorf("%s", "User is already logged in")
		h.Log.LogWarning(ctx, "users", "Login", err.Error())
		ctx.AbortWithStatus(http.StatusUnauthorized) // 401
		return
	}

	userSessionID, err := h.sessions.Create(loginData.Username, _const.CookieExpires)
	if err != nil {
		h.Log.LogError(ctx, "users", "Login", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	ctx.SetCookie(
		"session_id",
		userSessionID,
		int(_const.CookieExpires),
		"/",
		_const.Host,
		false,
		true,
	)

	ctx.Status(http.StatusOK) // 200
}

func (h *Handler) GetUser(ctx *gin.Context) {
	user, ok := ctx.Get(_const.UserKey)
	if !ok {
		err := fmt.Errorf("%s", "Failed to retrieve user from context")
		h.Log.LogWarning(ctx, "users", "GetUser", err.Error())
		ctx.AbortWithStatus(http.StatusNotFound) // 404
		return
	}

	userModel, ok := user.(models.User)
	if !ok {
		err := fmt.Errorf("%s", "Failed to cast user to model")
		h.Log.LogError(ctx, "users", "GetUser", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	userNoPassword := models.FromUser(userModel)
	ctx.JSON(http.StatusOK, userNoPassword)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	changed := new(models.User)
	err := ctx.BindJSON(changed)
	if err != nil {
		msg := "Failed to bind request data " + err.Error()
		h.Log.LogWarning(ctx, "users", "UpdateUser", msg)
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	user, ok := ctx.Get(_const.UserKey)
	if !ok {
		err := fmt.Errorf("%s", "Failed to retrieve user from context")
		h.Log.LogError(ctx, "users", "UpdateUser", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	userModel, ok := user.(models.User)
	if !ok {
		err := fmt.Errorf("%s", "Failed to cast user to model")
		h.Log.LogError(ctx, "users", "UpdateUser", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	changed.Username = userModel.Username
	changed.Avatar = userModel.Avatar
	newUser, err := h.useCase.UpdateUser(&userModel, *changed)
	if err != nil {
		h.Log.LogError(ctx, "users", "UpdateUser", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	userNoPassword := models.FromUser(*newUser)
	ctx.JSON(http.StatusOK, userNoPassword)
}

func (h *Handler) UploadAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		msg := "Failed to form file " + err.Error()
		h.Log.LogWarning(ctx, "users", "UploadAvatar", msg)
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	extension := filepath.Ext(file.Filename)
	// generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	err = ctx.SaveUploadedFile(file, _const.AvatarsFileDir+newFileName)

	if err != nil {
		h.Log.LogError(ctx, "users", "UploadAvatar", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	user, ok := ctx.Get(_const.UserKey)
	if !ok {
		err := fmt.Errorf("%s", "Failed to retrieve user from context")
		h.Log.LogError(ctx, "users", "UploadAvatar", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	userModel, ok := user.(models.User)
	if !ok {
		err := fmt.Errorf("%s", "Failed to cast user to model")
		h.Log.LogError(ctx, "users", "UploadAvatar", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	change := models.User{
		Username: userModel.Username,
		Avatar:   _const.AvatarsPath + newFileName,
	}
	//change.Avatar = _const.AvatarsPath + newFileName

	newUser, err := h.useCase.UpdateUser(&userModel, change)
	if err != nil {
		h.Log.LogError(ctx, "users", "UploadAvatar", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	userNoPassword := models.FromUser(*newUser)
	ctx.JSON(http.StatusOK, userNoPassword)
}
