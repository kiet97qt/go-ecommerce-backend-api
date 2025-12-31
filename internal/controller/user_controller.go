package controller

import (
	"net/http"

	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/pkg/loggers"
	"go-ecommerce-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger = loggers.GetLogger()

// UserController provides handlers for user resources.
type UserController struct {
	userService service.IUserService
}

// NewUserController binds a user controller to a service.
func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

type RegisterUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// GetUserByID returns a user by its identifier, responding with mock data.
func (uc *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := uc.userService.GetUserByID(ctx.Request.Context(), id)
	logger.Info("GetUserByID", zap.Any("user", user))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		response.ErrorResponse(ctx, 50001)
		return
	}

	response.SuccessResponse(ctx, user, 20001)
}

// RegisterUser nhận email, tạo OTP và gửi email OTP.
func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var req RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.userService.RegisterUser(ctx.Request.Context(), req.Email); err != nil {
		logger.Error("RegisterUser failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OTP sent to email",
	})
}
