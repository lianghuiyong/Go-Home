package function

import (
	"../data"
	"github.com/labstack/echo"
	"net/http"
	"github.com/astaxie/beego/orm"
)

//noinspection GoUnusedExportedFunction
func GetStations(c echo.Context) error{

	listStations := []data.StationBean{}

	o := orm.NewOrm()
	qs := o.QueryTable("station_bean")
	var posts []*data.StationBean
	qs.Limit(10000, 0).All(&posts)

	for _, val := range posts {
		listStations = append(listStations, *val)
	}

	ret := data.BaseResponse{http.StatusOK, "success", listStations}

	return c.JSONPretty(http.StatusOK, ret,"    ")
}

//noinspection ALL
func GetLeftTicket(c echo.Context) error{

	//解析Post参数
	train_date := c.Request().FormValue("train_date")
	from_station := c.Request().FormValue("from_station")
	to_station := c.Request().FormValue("to_station")

	LeftTicket(train_date,from_station,to_station)

	ret := data.BaseResponse{http.StatusOK, "success",LeftTicket(train_date,from_station,to_station) }

	return c.JSONPretty(http.StatusOK, ret,"    ")
}