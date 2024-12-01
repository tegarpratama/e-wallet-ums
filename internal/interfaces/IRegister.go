package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type IRegisterHandler interface {
	Register(c *gin.Context)
}

type RegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
