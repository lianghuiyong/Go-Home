package db

import (
	"github.com/astaxie/beego/config"
	"fmt"
	"github.com/astaxie/beego/orm"
)

var (
	dbType       string
	user         string
	password     string
	host         string
	port         string
	databaseName string
	dns          string
	conf, _      = config.NewConfig("ini", "db.config")
)

//noinspection GoUnusedExportedFunction
func SqlInit()  {
	dbType = conf.String("type")
	user = conf.String("user")
	password = conf.String("pass")
	host = conf.String("host")
	port = conf.String("port")
	databaseName = conf.String("databaseName")

	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, databaseName)

	// auto create database
	//createDb()

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dns)
	orm.RunSyncdb("default", false, true)
}
