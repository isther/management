package service

import (
	"github.com/isther/management/dao"
	"github.com/isther/management/model"
)

type AdminService struct{}

func NewAdminService() *AdminService { return &AdminService{} }

func (service *AdminService) Add(name string, i int) error {
	if name == "" {
		return nil
	}

	switch i {
	case 1:
		res := dao.DB.Create(&model.InBoundPersons{Name: name})
		return res.Error
	case 2:
		res := dao.DB.Create(&model.OutBoundPersons{Name: name})
		return res.Error
	case 3:
		res := dao.DB.Create(&model.Units{Name: name})
		return res.Error
	case 4:
		res := dao.DB.Create(&model.StrongLocation{Name: name})
		return res.Error
	}
	return nil
}

func (service *AdminService) Delete(name string, i int) error {
	switch i {
	case 1:
		res := dao.DB.Where("name = ?", name).Unscoped().Delete(&model.InBoundPersons{})
		return res.Error
	case 2:
		res := dao.DB.Where("name = ?", name).Unscoped().Delete(&model.OutBoundPersons{})
		return res.Error
	case 3:
		res := dao.DB.Where("name = ?", name).Unscoped().Delete(&model.Units{})
		return res.Error
	case 4:
		res := dao.DB.Where("name = ?", name).Unscoped().Delete(&model.StrongLocation{})
		return res.Error
	}
	return nil
}

func (service *AdminService) GetAll() (*model.AdminConfig, error) {
	var (
		inboundPersons  []model.InBoundPersons
		outboundPersons []model.OutBoundPersons
		units           []model.Units
		strongLocation  []model.StrongLocation
		itemIDs         []model.ItemIDs

		config model.AdminConfig
	)

	if res := dao.DB.Find(&inboundPersons); res.Error != nil {
		return &config, res.Error
	}
	config.InBoundPersons = inboundPersons

	if res := dao.DB.Find(&outboundPersons); res.Error != nil {
		return &config, res.Error
	}
	config.OutBoundPersons = outboundPersons

	if res := dao.DB.Find(&units); res.Error != nil {
		return &config, res.Error
	}
	config.Units = units

	if res := dao.DB.Find(&strongLocation); res.Error != nil {
		return &config, res.Error
	}
	config.StrongLocation = strongLocation

	if res := dao.DB.Find(&itemIDs); res.Error != nil {
		return &config, res.Error
	}
	config.ItemIDs = itemIDs

	return &config, nil
}
