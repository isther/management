package service

import (
	"github.com/isther/management/dao"
	"github.com/isther/management/model"
	"gorm.io/gorm"
)

type ItemService struct{}

func NewItemService() *ItemService { return &ItemService{} }

func (service *ItemService) CreateItemAndItemID(item model.Item) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.DB.Create(&model.ItemSql{Item: item}).Error; err != nil {
			return err
		}
		return dao.DB.Create(&model.ItemIDs{Name: item.ItemID}).Error
	})
}

func (service *ItemService) QueryAll() ([]model.ItemSql, error) {
	var (
		itemSqls []model.ItemSql
	)
	if res := dao.DB.Find(&itemSqls); res.Error != nil {
		return nil, res.Error
	}

	return itemSqls, nil
}

func (service *ItemService) Update(newItem model.Item) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Where("item_id = ?", newItem.ItemID).Updates(&model.ItemSql{Item: newItem}).Error
	})
}

func (service *ItemService) DeleteByIDAndDeleteDetails(id string) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		var itemSql model.ItemSql
		if err := tx.Where("id = ?", id).First(&itemSql).Error; err != nil {
			return err
		}

		if err := tx.Where("item_id = ?", itemSql.ItemID).Unscoped().Delete(&model.InboundSql{}).Error; err != nil {
			return err
		}

		if err := tx.Where("item_id = ?", itemSql.ItemID).Unscoped().Delete(&model.OutboundSql{}).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", id).Unscoped().Delete(&model.ItemSql{}).Error; err != nil {
			return err
		}

		return tx.Where("name = ?", itemSql.ItemID).Unscoped().Delete(&model.ItemIDs{}).Error
	})
}

func (service *ItemService) DeleteByItemIDAndDeleteDetails(item_id string) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("item_id = ?", item_id).Unscoped().Delete(&model.InboundSql{}).Error; err != nil {
			return err
		}

		if err := tx.Where("item_id = ?", item_id).Unscoped().Delete(&model.OutboundSql{}).Error; err != nil {
			return err
		}

		return tx.Where("item_id = ?", item_id).Unscoped().Delete(&model.ItemSql{}).Error
	})
}
