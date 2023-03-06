package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/isther/management/model"
	"github.com/isther/management/service"
	"github.com/sirupsen/logrus"
)

type ItemApi struct{}

func NewItemApi() *ItemApi { return &ItemApi{} }

func (api *ItemApi) Create(ctx *gin.Context) {
	var (
		item model.Item
	)
	if err := ctx.ShouldBindJSON(&item); err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := service.NewItemService().CreateItemAndItemID(item); err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := service.NewAdminService().Add(item.ItemID, 5); err != nil {
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

func (api *ItemApi) Query(ctx *gin.Context) {
	items, err := service.NewItemService().QueryAll()
	if err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg":   "ok",
		"items": items,
	})
}

func (api *ItemApi) Update(ctx *gin.Context) {
	var (
		item model.Item
	)
	if err := ctx.ShouldBindJSON(&item); err != nil {
		logrus.Error(err)
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := service.NewItemService().Update(item); err != nil {
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

func (api *ItemApi) Delete(ctx *gin.Context) {
	id := ctx.Query("id")

	if err := service.NewItemService().DeleteByIDAndDeleteDetails(id); err != nil {
		ctx.JSON(STATUS_FAILED, gin.H{
			"msg": err,
		})
		return
	}

	ctx.JSON(STATUS_SUCCESS, gin.H{
		"msg": "ok",
	})
}
