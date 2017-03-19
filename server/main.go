package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"./api"
	"./db"
	"./data"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbType       string
	user         string
	password     string
	host         string
	port         string
	databaseName string
)

func init() {
	orm.Debug = true

	var conf ,err = config.NewConfig("ini", "./db/db.config")
	if err != nil{
		fmt.Println(err)
	}

	dbType = conf.String("type")
	user = conf.String("user")
	password = conf.String("pass")
	host = conf.String("host")
	port = conf.String("port")
	databaseName = conf.String("databaseName")

	//dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, databaseName)
	dataSource := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",user, password,databaseName, host, port)

	fmt.Println(dataSource)

	orm.RegisterDriver(dbType, orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres",  dataSource)

	// 创建 table
	orm.RunSyncdb("default", false, true)

	// 需要在init中注册定义的model
	orm.RegisterModel(new(data.StationBean))
}

func initGoHome()  {
	db.InitStations()
}

func startEcho()  {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//static file serviing
	e.Static("/static", "assets")

	// Routers
	e.POST("/getstations", api.GetStationNames)
	//e.GET("/users/:id", controllers.ShowUser)
	//e.GET("/users", controllers.AllUsers)
	//e.PUT("/users/:id", controllers.UpdateUser)
	//e.DELETE("/users/:id",controllers.DeleteUser)

	// Server
	e.Start(":1323")
}

func main() {

	fmt.Println("数据库初始化...")
	initGoHome()

	fmt.Println("正在开启服务器...")
	startEcho()
}
