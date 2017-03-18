package db

import (
	"io/ioutil"
	"strings"
	"net/http"
	"../api"
	"fmt"
)

// 初始化车站信息
func InitStations()  {
	resp, _ := http.Get(api.StationNameURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	timBody1 := strings.Trim(string(body),"var station_names ='")
	timBody2 := strings.Trim(timBody1,"';\n")

	split1 := strings.Split(timBody2,"@")

	for index, value := range split1 {
		fmt.Printf("arr[%d]=%s \n", index, value)
	}
}
