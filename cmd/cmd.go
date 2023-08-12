package cmd

import (
	"fmt"

	"github.com/ikun666/go_webserver/conf"
	"github.com/ikun666/go_webserver/dao"
	"github.com/ikun666/go_webserver/global"
	"github.com/ikun666/go_webserver/router"
)

func Start() {
	//初始化配置文件
	conf.InitConfig()
	//初始化数据库
	db, err := dao.InitDB()
	if err != nil {
		fmt.Println("init db err")
	}
	//将数据库保存到全局变量
	global.DB = db
	fmt.Println("init db success")
	//初始化路由
	router.InitRouter()
}
func End() {
	fmt.Println("end")
}
