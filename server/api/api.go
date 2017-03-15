package api

import (
	"../data"
	"github.com/labstack/echo"
	"net/http"
	"io/ioutil"
	"strings"
)

//noinspection GoUnusedExportedFunction
func PostTest(c echo.Context) error{

	movies := []data.MovieBean{
		{"金刚狼", "2017", []string{"休·杰克曼", "达芙妮·基恩", "帕特里克·斯图尔特"}},
		{"极限特工", "2017", []string{"范·迪塞尔", "露比·罗丝", "妮娜·杜波夫"}},
		{"功夫瑜伽", "2017", []string{"成龙", "李治廷", "张艺兴"}},
		{"生化危机：终章", "2017", []string{"米拉·乔沃维奇", "伊恩·格雷", "艾丽·拉特"}},
	}

	baseMovie := data.BaseResponse{http.StatusOK, "success", movies}

	return c.JSONPretty(http.StatusOK, baseMovie,"    ")
}



//noinspection GoUnusedExportedFunction
func GetStationNames(c echo.Context) error{
	resp, _ := http.Get(StationNameURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	timBody1 := strings.Trim(string(body),"var station_names ='")
	timBody2 := strings.Trim(timBody1,"';\n")

	baseMovie := data.BaseResponse{http.StatusOK, "success", timBody2}

	return c.JSONPretty(http.StatusOK, baseMovie,"    ")
}