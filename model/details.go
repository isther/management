package model

import "gorm.io/gorm"

type InboundSql struct {
	gorm.Model
	Details
}

type OutboundSql struct {
	gorm.Model
	Details
}

// details
type Details struct {
	Timestamp string `json:"timestamp"` // 时间戳
	ItemID    string `json:"item_id"`   // 物料编号
	// ItemName       string `json:"item_name"`       // 物料名称
	// Specification  string `json:"specification"`   // 规格型号
	// Unit           string `json:"unit"`            // 计量单位
	// StrongLocation string `json:"strong_location"` // 库位
	Number  string `json:"number"`  // 数量
	Person  string `json:"person"`  // 申请人
	Comment string `json:"comment"` // 备注
}
