package service

import (
	"github.com/isther/management/dao"
	"github.com/isther/management/model"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type DetailsService struct{}

func NewDetailsService() *DetailsService { return &DetailsService{} }

func (service *DetailsService) CreateInbound(detail model.Details) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		var itemSql model.ItemSql
		if err := tx.Where("item_id = ?", detail.ItemID).First(&itemSql).Error; err != nil {
			return err
		}

		detailNumber, err := decimal.NewFromString(detail.Number)
		if err != nil {
			return err
		}

		itemNumber, err := decimal.NewFromString(itemSql.Number)
		if err != nil {
			return err
		}

		itemInNumber, err := decimal.NewFromString(itemSql.InNumber)
		if err != nil {
			return err
		}

		itemSql.InNumber = itemInNumber.Add(detailNumber).String()
		itemSql.Number = itemNumber.Add(detailNumber).String()

		if err := tx.Where("item_id = ?", detail.ItemID).Updates(&itemSql).Error; err != nil {
			return err
		}

		return tx.Create(&model.InboundSql{Details: detail}).Error
	})
}

func (service *DetailsService) CreateOutbound(detail model.Details) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		var itemSql model.ItemSql
		if err := tx.Where("item_id = ?", detail.ItemID).First(&itemSql).Error; err != nil {
			return err
		}

		detailNumber, err := decimal.NewFromString(detail.Number)
		if err != nil {
			return err
		}

		itemNumber, err := decimal.NewFromString(itemSql.Number)
		if err != nil {
			return err
		}

		itemOutNumber, err := decimal.NewFromString(itemSql.OutNumber)
		if err != nil {
			return err
		}

		itemSql.OutNumber = itemOutNumber.Add(detailNumber).String()
		itemSql.Number = itemNumber.Sub(detailNumber).String()

		if err := tx.Where("item_id = ?", detail.ItemID).Updates(&itemSql).Error; err != nil {
			return err
		}

		return tx.Create(&model.OutboundSql{Details: detail}).Error
	})
}

type Detail struct {
	ID             string `json:"id"`
	Timestamp      string `json:"timestamp"`       // 时间戳
	ItemID         string `json:"item_id"`         // 物料编号
	ItemName       string `json:"item_name"`       // 物料名称
	Specification  string `json:"specification"`   // 规格型号
	Unit           string `json:"unit"`            // 计量单位
	StrongLocation string `json:"strong_location"` // 库位
	Number         string `json:"number"`          // 数量
	Person         string `json:"person"`          // 申请人
	Comment        string `json:"comment"`         // 备注
}

func (service *DetailsService) QueryAllInbound() ([]Detail, error) {
	var details []Detail
	/*	SELECT inbound_sqls.*,item_sqls.item_id,item_sqls.item_name,item_sqls.specification,item_sqls.unit,item_sqls.storage_location
			FROM inbound_sqls
		 	JOIN item_sqls
				ON inbound_sqls.item_id = item_sqls.item_id;*/
	dao.DB.Model(&model.InboundSql{}).
		Select("inbound_sqls.*,item_sqls.item_id,item_sqls.item_name,item_sqls.specification,item_sqls.unit,item_sqls.storage_location").
		Joins("JOIN item_sqls ON inbound_sqls.item_id = item_sqls.item_id").Scan(&details)

	return details, nil
}

func (service *DetailsService) QueryAllOutbound() ([]Detail, error) {
	var details []Detail
	/*	SELECT outbound_sqls.*,item_sqls.item_id,item_sqls.item_name,item_sqls.specification,item_sqls.unit,item_sqls.storage_location
			FROM outbound_sqls
		 	JOIN item_sqls
				ON outbound_sqls.item_id = item_sqls.item_id;*/
	dao.DB.Model(&model.OutboundSql{}).
		Select("outbound_sqls.*,item_sqls.item_id,item_sqls.item_name,item_sqls.specification,item_sqls.unit,item_sqls.storage_location").
		Joins("JOIN item_sqls ON outbound_sqls.item_id = item_sqls.item_id").Scan(&details)

	return details, nil
}

func (service *DetailsService) QueryByTimestamp(start, end string) ([]Detail, []Detail, error) {
	var (
		inBoundDetails  []Detail
		outBoundDetails []Detail
	)

	dao.DB.Model(&model.InboundSql{}).
		Select("inbound_sqls.*,item_sqls.item_id,item_sqls.item_name,item_sqls.specification,item_sqls.unit,item_sqls.storage_location").
		Joins("JOIN item_sqls ON inbound_sqls.item_id = item_sqls.item_id").
		Where("timestamp >= ?", start).
		Where("timestamp <= ?", end).
		Scan(&inBoundDetails)

	dao.DB.Model(&model.OutboundSql{}).
		Select("outbound_sqls.*,item_sqls.item_id,item_sqls.item_name,item_sqls.specification,item_sqls.unit,item_sqls.storage_location").
		Joins("JOIN item_sqls ON outbound_sqls.item_id = item_sqls.item_id").
		Where("timestamp >= ?", start).
		Where("timestamp <= ?", end).
		Scan(&outBoundDetails)

	return inBoundDetails, outBoundDetails, nil
}

func (service *DetailsService) UpdateInboundByID(id string) error {
	return dao.DB.Where("id = ?", id).Unscoped().Delete(&model.InboundSql{}).Error
}

func (service *DetailsService) UpdateOutBoundByID(id string) error {
	return dao.DB.Where("id = ?", id).Unscoped().Delete(&model.OutboundSql{}).Error
}

func (service *DetailsService) DeleteInboundByID(id string) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		var (
			inboundSql model.InboundSql
			itemSql    model.ItemSql
		)
		if err := tx.Where("id = ?", id).First(&inboundSql).Error; err != nil {
			return err
		}

		if err := tx.Where("item_id = ?", inboundSql.ItemID).First(&itemSql).Error; err != nil {
			return err
		}

		detailNumber, err := decimal.NewFromString(inboundSql.Number)
		if err != nil {
			return err
		}

		itemNumber, err := decimal.NewFromString(itemSql.Number)
		if err != nil {
			return err
		}

		itemInNumber, err := decimal.NewFromString(itemSql.InNumber)
		if err != nil {
			return err
		}

		itemSql.InNumber = itemInNumber.Sub(detailNumber).String()
		itemSql.Number = itemNumber.Sub(detailNumber).String()

		if err := tx.Where("item_id = ?", itemSql.ItemID).Updates(&itemSql).Error; err != nil {
			return err
		}

		return tx.Where("id = ?", id).Unscoped().Delete(&model.InboundSql{}).Error
	})
}

func (service *DetailsService) DeleteOutBoundByID(id string) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		var (
			outboundSql model.OutboundSql
			itemSql     model.ItemSql
		)
		if err := tx.Where("id = ?", id).First(&outboundSql).Error; err != nil {
			return err
		}

		if err := tx.Where("item_id = ?", outboundSql.ItemID).First(&itemSql).Error; err != nil {
			return err
		}

		detailNumber, err := decimal.NewFromString(outboundSql.Number)
		if err != nil {
			return err
		}

		itemNumber, err := decimal.NewFromString(itemSql.Number)
		if err != nil {
			return err
		}

		itemInNumber, err := decimal.NewFromString(itemSql.InNumber)
		if err != nil {
			return err
		}

		itemSql.OutNumber = itemInNumber.Sub(detailNumber).String()
		itemSql.Number = itemNumber.Add(detailNumber).String()

		if err := tx.Where("item_id = ?", itemSql.ItemID).Updates(&itemSql).Error; err != nil {
			return err
		}

		return tx.Where("id = ?", id).Unscoped().Delete(&model.OutboundSql{}).Error
	})
}
