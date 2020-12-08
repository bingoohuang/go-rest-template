package api

import (
	"fmt"

	"github.com/bingoohuang/go-rest-template/pkg/helpers"

	"github.com/bingoohuang/go-rest-template/internal/api/router"
	"github.com/bingoohuang/go-rest-template/internal/pkg/conf"
	"github.com/bingoohuang/go-rest-template/internal/pkg/db"
	"github.com/gin-gonic/gin"
)

func setConfiguration(configPath string) {
	conf.Setup(configPath)
	db.SetupDB()
	gin.SetMode(conf.GetConf().Server.Mode)
}

func Run(configPath string) {
	setConfiguration(helpers.DefaultTo(configPath, "data/conf.yml"))
	c := conf.GetConf()
	web := router.Setup()

	fmt.Println("REST API Running on port", c.Server.Port)
	_ = web.Run(fmt.Sprintf(":%d", c.Server.Port))
}
