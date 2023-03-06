package dao

import (
	"github.com/isther/management/conf"
	"github.com/isther/management/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	initSQL()
}

func initSQL() {
	var err error

	// 连接数据库
	DB, err = gorm.Open(postgres.Open(conf.Server.DSN()), &gorm.Config{})
	if err != nil {
		logrus.Fatalln(err)
	}

	// 绑定模型
	err = DB.AutoMigrate(&model.UserSql{}, &model.ItemSql{}, &model.InboundSql{}, &model.OutboundSql{})
	if err != nil {
		logrus.Fatalln(err)
	}

	err = DB.AutoMigrate(&model.InBoundPersons{}, &model.OutBoundPersons{}, &model.Units{}, &model.StrongLocation{}, &model.ItemIDs{})
	if err != nil {
		logrus.Fatalln(err)
	}
}
