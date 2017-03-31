package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"./api"
	"./db"
	"./function"

	//导入MySQL驱动
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func initGoHome() {
	db.InitMySQL()
	function.InitStations()
}

func startEcho() {
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
	e.POST("/stations", api.GetStations)
	e.POST("/test", api.Test)
	//e.GET("/users/:id", controllers.ShowUser)
	//e.GET("/users", controllers.AllUsers)
	//e.PUT("/users/:id", controllers.UpdateUser)
	//e.DELETE("/users/:id",controllers.DeleteUser)

	// Server
	e.Start(":1323")
}

func main() {

	fmt.Println(">>>>数据库初始化...")
	initGoHome()

	fmt.Println(">>>>正在开启服务器...")
	startEcho()
}
