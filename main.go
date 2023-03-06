package main

import (
	"github.com/isther/management/conf"
	"github.com/isther/management/routers"
	"github.com/sirupsen/logrus"
)

func main() {
	r := routers.Init()
	logrus.Info("Server listen: ", conf.Server.Listen)
	r.Run(conf.Server.Listen)
}
