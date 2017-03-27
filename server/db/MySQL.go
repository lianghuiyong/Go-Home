package db

import (
	"../data"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"

	//导入MySQL驱动
	_ "github.com/go-sql-driver/mysql"
)

var (
	user         string
	password     string
	host         string
	port         string
	databaseName string
)

func InitMySQL()  {
	var conf, err = config.NewConfig("ini", "./db/db.config")
	if err != nil {
		fmt.Println(err)
	}

	//dbType = conf.String("type")
	user = conf.String("user")
	password = conf.String("pass")
	host = conf.String("host")
	port = conf.String("port")
	databaseName = conf.String("databaseName")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, databaseName)

	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 设置默认数据库
	//mysql用户：root ，密码：zxxx ， 数据库名称：test ， 数据库别名：default
	orm.RegisterDataBase("default", "mysql", dataSource)

	// 注册定义的 model
	orm.RegisterModel(new(data.StationBean))
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))

	// 创建 table
	orm.RunSyncdb("default", false, true)
}
