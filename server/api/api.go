package api

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
	qs.Limit(3000, 0).All(&posts)

	for _, val := range posts {
		listStations = append(listStations, *val)
	}

	baseMovie := data.BaseResponse{http.StatusOK, "success", listStations}

	return c.JSONPretty(http.StatusOK, baseMovie,"    ")
}

//noinspection GoUnusedExportedFunction
func Test(c echo.Context) error{


	baseMovie := data.BaseResponse{http.StatusOK, "success", "hello"}

	return c.JSONPretty(http.StatusOK, baseMovie,"    ")
}