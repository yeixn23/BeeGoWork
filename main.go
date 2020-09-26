package main

import (
	_ "BeegoHello/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//appName := config.String("appname")
	//fmt.Println("应用名称",appName)
	//port,err := config.Int("httpport")
	//if err != nil {
	//	panic("项目配置文件解析失败，请检查配置文件")
	//}

	beego.Run()
}
/*
 *git add .
*git commit -m "注释"
*git push
 */
