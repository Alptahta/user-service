package rest

import (
	"github.com/Alptahta/user-service/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		svc: userService,
	}
}

func (u *UserHandler) Register(r *gin.Engine) {

}
