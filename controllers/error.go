package controllers

import "github.com/astaxie/beego"
/**
  该控制器处理页面错误请求
 */
type ErrorController struct {
	beego.Controller
}
func (c *ErrorController) Error401() {
	c.Data["json"] = map[string]interface{}{ "code":"401","err" :"未经授权，请求要求验证身份" }
	c.ServeJSON()
}
func (c *ErrorController) Error403() {
	c.Data["json"] =  map[string]interface{}{ "code":"403","err" :"服务器拒绝请求" }
	c.ServeJSON()
}
func (c *ErrorController) Error404() {
	c.Data["json"] =  map[string]interface{}{ "code":"404","err" :"很抱歉您访问的地址或者方法不存在" }
	c.ServeJSON()
}

func (c *ErrorController) Error500() {
	c.Data["json"] =  map[string]interface{}{ "code":"500","err" :"服务器错误" } 
	c.ServeJSON()
}
func (c *ErrorController) Error503() {
	c.Data["json"] =  map[string]interface{}{ "code":"404","err" :"服务器目前无法使用（由于超载或停机维护）" }
	c.ServeJSON()
}
func (c *ErrorController) Error505(){

}
