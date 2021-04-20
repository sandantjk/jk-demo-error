package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sandantjk/jk-demo-error/internal/conf"
	"github.com/sandantjk/jk-demo-error/internal/service"
)

const configFile = "./configs/application.yml"

func main() {
	var id int64
	flag.Int64Var(&id, "i", id, "用户ID")
	flag.Parse()
	initApp()
	defer closeApp()
	// 业务处理
	err := service.PrintUserInfo(id)
	// 框架统一处理异常
	if err != nil {
		fmt.Println(fmt.Sprintf("访问数据库发上错误:\n%+v", err))
	}
}

// 初始化
func initApp() {
	config, err := conf.InitConfig(configFile)
	if err != nil {
		panic(err)
	}
	db, err := conf.InitDB(config.DB)
	if err != nil {
		panic(err)
	}
	conf.Global = config
	conf.DB = db
}

// 关闭
func closeApp() {
	if conf.DB != nil {
		_ = conf.DB.Close()
	}
}
