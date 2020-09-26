package controllers

import (
	"BeegoHello/db_mysql"
	"BeegoHello/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

/*
 *该结构体用于处理/register
 */
type RegisterController struct {
	beego.Controller
}
/*
 *该方法用于处理post类型请求
 */
func (r *RegisterController) Post() {
	fmt.Println(r == nil)
	fmt.Println(r.Ctx == nil)
	fmt.Println(r.Ctx.Request == nil)
	//1.接受前端传递来的数据
	//r.Ctx.Request.GetBody()
	//if err !=nil {
	//	r.Ctx.WriteString("数据接受错误")
	//	return
	//}
	//
	bodyBytes,err :=ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("数据接受错误，请重试")
		return
	}
	var user models.User
	err = json.Unmarshal(bodyBytes,&user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据解析错误")
		return
	}
	//将解析到的用户，保存到数据
	id, err := db_mysql.InserUser(user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户保存失败.")
		return
	}
	fmt.Println(id)
	result := models.ResponseResult{
		Code:    0,
		Message: "保存成功",
		Date:    nil,
	}
	r.Data["json"] = &result
	r.ServeJSON()
}
