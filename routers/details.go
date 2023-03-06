package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/isther/management/model"
	"github.com/isther/management/service"
	"github.com/sirupsen/logrus"
)

type DetailsApi struct{}

func NewDetailsApi() *DetailsApi { return &DetailsApi{} }

func (api *DetailsApi) CreateInbound(ctx *gin.Context) {
	var (
		detail model.Details
	)
	if err := ctx.ShouldBindJSON(&detail); err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := service.NewDetailsService().CreateInbound(detail); err != nil {
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

func (api *DetailsApi) QueryAllInbound(ctx *gin.Context) {
	details, err := service.NewDetailsService().QueryAllInbound()
	if err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg":     "ok",
		"details": details,
	})
}

func (api *DetailsApi) CreateOutbound(ctx *gin.Context) {
	var (
		detail model.Details
	)
	if err := ctx.ShouldBindJSON(&detail); err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := service.NewDetailsService().CreateOutbound(detail); err != nil {
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

func (api *DetailsApi) QueryAllOutbound(ctx *gin.Context) {
	details, err := service.NewDetailsService().QueryAllOutbound()
	if err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg":     "ok",
		"details": details,
	})
}

func (api *DetailsApi) QueryByTimestamp(ctx *gin.Context) {
	start := ctx.Query("start")
	end := ctx.Query("end")

	inBoundDetails, outBoundDetails, err := service.NewDetailsService().QueryByTimestamp(start, end)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg":      "ok",
		"inbound":  inBoundDetails,
		"outbound": outBoundDetails,
	})
}

func (api *DetailsApi) UpdateInbound(ctx *gin.Context) {
	id := ctx.Query("id")

	if err := service.NewDetailsService().UpdateInboundByID(id); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg": "ok",
	})
}

func (api *DetailsApi) UpdateOutbound(ctx *gin.Context) {
	id := ctx.Query("id")

	if err := service.NewDetailsService().UpdateOutBoundByID(id); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg": "ok",
	})
}

func (api *DetailsApi) DeleteInbound(ctx *gin.Context) {
	id := ctx.Query("id")

	if err := service.NewDetailsService().DeleteInboundByID(id); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg": "ok",
	})
}

func (api *DetailsApi) DeleteOutbound(ctx *gin.Context) {
	id := ctx.Query("id")

	if err := service.NewDetailsService().DeleteOutBoundByID(id); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg": "ok",
	})
}
