package db

import (
	"io/ioutil"
	"strings"
	"net/http"
	"../api"
	"../data"
	"fmt"
	"github.com/labstack/gommon/log"
	"crypto/tls"
	"github.com/astaxie/beego/orm"
)

// 初始化车站信息
//noinspection GoUnusedExportedFunction
func InitStations() {
	client := newClient()

	resp, err := client.Get(api.StationNameURL)
	if err != nil {
		log.Fatal(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(">>>> 更新数据库车站信息...")

	timBody1 := strings.Trim(string(body), "var station_names ='")
	timBody2 := strings.Trim(timBody1, "';\n")

	split1 := strings.Split(timBody2, "@")

	//获取orm
	o := orm.NewOrm()

	//关闭写同步
	o.Raw("PRAGMA synchronous = OFF; ", 0, 0, 0).Exec()
	fmt.Println(">>>> beego Orm 关闭写同步！")

	//清空表
	o.Raw("TRUNCATE TABLE station_bean").Exec()
	fmt.Println(">>>> station_bean 表已清空！")

	for _, value := range split1 {

		split2 := strings.Split(value, "|")
		if len(split2) > 1 {
			station := data.StationBean{}
			for i := 0; i < len(split2); i++ {
				switch i {
				case 0:
					station.Namejp = split2[i]
					break
				case 1:
					station.Name = split2[i]
					break
				case 2:
					station.Namewz = split2[i]
					break
				case 3:
					station.Nameqp = split2[i]
					break
				case 4:
					station.Namejp2 = split2[i]
					break
				case 5:
					station.Bianhao = split2[i]
					break
				}
			}

			// 插入表
			o.Insert(&station)
		}
	}
	fmt.Println(">>>> 更新数据库车站完成！")
}

//解决 x509 未认证的验证问题
func newClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}
