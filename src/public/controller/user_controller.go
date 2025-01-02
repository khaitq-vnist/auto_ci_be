package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/request"
	"github.com/khaitq-vnist/auto_ci_be/public/apihelper"
	request2 "github.com/khaitq-vnist/auto_ci_be/public/resource/request"
	"github.com/khaitq-vnist/auto_ci_be/public/resource/response"
	"github.com/khaitq-vnist/auto_ci_be/public/service"
)

type UserController struct {
	userService service.IUserService
}

func (u *UserController) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Error while binding request", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	result, err := u.userService.CreateUser(context.Background(), &req)
	if err != nil {
		log.Error(c, "Error while creating user", err)
		if err.Error() == common.ExistedEmailMessage {
			apihelper.AbortErrorHandle(c, common.ExistedEmailCode)
			return
		}
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromEntityToUserResponse(result))
}
func (u *UserController) LoginUser(c *gin.Context) {
	var req request2.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Error while binding request", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	result, err := u.userService.LoginUser(context.Background(), req.Email, req.Password)
	if err != nil {
		log.Error(c, "Error while login user", err)
		if err.Error() == common.ErrInvalidPassword {
			apihelper.AbortErrorHandle(c, common.InvalidPasswordCode)
			return
		}

		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, result)
}
func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}
