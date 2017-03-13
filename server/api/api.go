package api

import (
	"../data"
	"github.com/labstack/echo"
	"net/http"
)

//noinspection GoUnusedExportedFunction
func PostTest(c echo.Context) error{

	movies := []data.Movie{
		{"金刚狼", "2017", []string{"休·杰克曼", "达芙妮·基恩", "帕特里克·斯图尔特"}},
		{"极限特工", "2017", []string{"范·迪塞尔", "露比·罗丝", "妮娜·杜波夫"}},
		{"功夫瑜伽", "2017", []string{"成龙", "李治廷", "张艺兴"}},
		{"生化危机：终章", "2017", []string{"米拉·乔沃维奇", "伊恩·格雷", "艾丽·拉特"}},
	}

	baseMovie := data.BaseResponse{http.StatusOK, "success", movies}

	return c.JSONPretty(http.StatusOK, baseMovie,"  ")
}
