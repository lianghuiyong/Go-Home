package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"./api"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	e := echo.New()

	//启动时获取一次数据到本地数据库
	resp, _ := http.Get(api.StationNameURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

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
	e.POST("/stations", api.PostTest)
	//e.GET("/users/:id", controllers.ShowUser)
	//e.GET("/users", controllers.AllUsers)
	//e.PUT("/users/:id", controllers.UpdateUser)
	//e.DELETE("/users/:id",controllers.DeleteUser)

	// Server
	e.Start(":1323")
}
