package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isther/management/middleware"
	"github.com/isther/management/service"
	"github.com/sirupsen/logrus"
)

type AdminApi struct{}

func NewAdminApi() *AdminApi { return &AdminApi{} }

func (api *AdminApi) Add(ctx *gin.Context) {
	if api.getUsername(ctx) != "admin" {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": "Not admin",
		})
		return
	}

	inbound_persons := ctx.Query("inbound_persons")
	outbound_persons := ctx.Query("outbound_persons")
	units := ctx.Query("units")
	strong_locations := ctx.Query("strong_locations")

	if err := service.NewAdminService().Add(inbound_persons, 1); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}
	if err := service.NewAdminService().Add(outbound_persons, 2); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}
	if err := service.NewAdminService().Add(units, 3); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}
	if err := service.NewAdminService().Add(strong_locations, 4); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}

	config, err := service.NewAdminService().GetAll()
	if err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg":    "ok",
		"config": config,
	})
}

func (api *AdminApi) Delete(ctx *gin.Context) {
	if api.getUsername(ctx) != "admin" {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": "Not admin",
		})
		return
	}

	inbound_persons := ctx.Query("inbound_persons")
	outbound_persons := ctx.Query("outbound_persons")
	units := ctx.Query("units")
	strong_locations := ctx.Query("strong_locations")

	if err := service.NewAdminService().Delete(inbound_persons, 1); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}
	if err := service.NewAdminService().Delete(outbound_persons, 2); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}
	if err := service.NewAdminService().Delete(units, 3); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}
	if err := service.NewAdminService().Delete(strong_locations, 4); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}

	config, err := service.NewAdminService().GetAll()
	if err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg":    "ok",
		"config": config,
	})
}

func (api *AdminApi) Get(ctx *gin.Context) {
	if api.getUsername(ctx) != "admin" {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": "Not admin",
		})
		return
	}

	config, err := service.NewAdminService().GetAll()
	if err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg":    "ok",
		"config": config,
	})
}

func (api *AdminApi) getUsername(ctx *gin.Context) string {
	token := ctx.GetHeader("Authorization")
	claims, err := middleware.ParseMapClaimsJwt(token)
	if err != nil {
		logrus.Error("failed to generate token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return ""
	}
	return claims["username"].(string)
}
