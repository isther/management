package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/isther/management/middleware"
	"github.com/isther/management/model"
	"github.com/isther/management/service"
	"github.com/sirupsen/logrus"
)

type UserApi struct{}

func NewUserApi() *UserApi { return &UserApi{} }

func (api *UserApi) Register(ctx *gin.Context) {
	var (
		user model.User
	)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := service.NewUserService().Create(user); err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg": "ok",
	})
}

func (api *UserApi) Login(ctx *gin.Context) {
	var (
		user model.User
	)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// Find user by username
	tempUser, err := service.NewUserService().FindUserByUsername(user.Username)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// Compare password
	if tempUser.Username != user.Username || tempUser.Password != user.Password {
		logrus.Error("Error password")
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": "Error password",
		})
		return
	}

	token, err := middleware.BuildMapClaimsJwt(user.Username, user.Password)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	config, err := service.NewAdminService().GetAll()
	if err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg":   err,
			"token": token,
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg":    "ok",
		"token":  token,
		"config": config,
	})
}
