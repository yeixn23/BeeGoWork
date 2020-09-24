package controllers

import (
	"BeegoHello/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//1、获取请求数据
	userName := c.Ctx.Input.Query("user")
	password := c.Ctx.Input.Query("psd")
	//2、使用固定数据进行数据校验
	if userName !="yexin" || password != "123456"{
		//代表错误处理
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("请求成功"))
	//c.Data["Website"] = "www.baidu.com"
	//c.Data["Email"] = "aaa2398546675@163.com"
	//c.TplName = "index.tpl"
}
//编写一个post方法，用于处理post类型请求
func (c *MainController) post()  {
	for i := 0; i < 10 ;i++  {
		fmt.Printf("第%d次打印：",i)
	}
}
//func (c *MainController) Post() {
//	//1.接收（解析）post请求的参数
//	name := c.Ctx.Request.FormValue("name")
//	age := c.Ctx.Request.FormValue("age")
//	sex := c.Ctx.Request.FormValue("sex")
//	fmt.Println(name,age,sex)
//	//2、进行数据校验
//	if name != "admin" && age != "18" {
//		c.Ctx.WriteString("数据校验失败")
//		return
//	}
//	c.Ctx.WriteString("数据请求成功")
//}
/*
	该方法用于处理post类型的请求
 */
func (c *MainController) Post() {
	//1,解析前端提交的json格式数据
	var person models.Person
	dataBytes,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收错误,请重试")
		return
	}
	err = json.Unmarshal(dataBytes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("姓名：",person.Name)
	fmt.Println("年龄：",person.Age)
	fmt.Println("性别：",person.Sex)
	c.Ctx.WriteString("数据请求成功")
}

