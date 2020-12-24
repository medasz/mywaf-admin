package router

import "gopkg.in/macaron.v1"

//返回错误信息
func ErrorMsg(ctx *macaron.Context,message string){
	println("----")
	ctx.Data["message"] = message
	ctx.HTML(200,"error")
}
