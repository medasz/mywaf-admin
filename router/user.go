package router

import (
	"github.com/go-macaron/csrf"
	"gopkg.in/macaron.v1"
	"log"
	"mywaf-admin/model"
	"mywaf-admin/utils"
)

func UserList(ctx *macaron.Context) {
	users, err := model.GetAllUser()
	if err != nil {
		log.Printf("Fail to find user :", err)
		ErrorMsg(ctx, "获取用户信息失败!")
	}
	ctx.Data["users"] = users
	ctx.Data["rulesType"] = model.RuleInfo
	ctx.HTML(200, "user")
}

func DeleteUser(ctx *macaron.Context) {
	id := ctx.ParamsInt64(":Id")
	err := model.DeleteUserById(id)
	if err != nil {
		log.Printf("Fail to delete user id:%v :%s", id, err)
		ErrorMsg(ctx, "删除用户失败!")
	}
	ctx.Redirect("/index/user/list/")
}

func EditUser(ctx *macaron.Context, x csrf.CSRF) {
	id := ctx.ParamsInt64(":Id")
	user, err := model.GetUserById(id)
	if err != nil {
		log.Printf("Fail to get user by id:%v :%s", id, err)
		ErrorMsg(ctx, "获取用户信息失败!")
	}
	ctx.Data["user"] = user
	ctx.Data["Id"] = id
	ctx.Data["csrf_token"] = x.GetToken()
	ctx.HTML(200, "editUser")
}

func EditUserDeal(ctx *macaron.Context) {
	id := ctx.ParamsInt64(":Id")
	username := ctx.Req.PostForm.Get("username")
	password := ctx.Req.PostForm.Get("password")
	if password != "" {
		password = utils.Encryption(password)
	}
	log.Printf("username:%s;password:%s", username, password)
	err := model.UpdateUser(id, username, password)
	if err != nil {
		log.Printf("Fail to update user{Id:%v,Username:%s,Password:%s} :%s", id, username, password, err)
		ErrorMsg(ctx, "修改用户信息失败!")
	}
	ctx.Redirect("/index/user/list/")
}

func AddUser(ctx *macaron.Context, x csrf.CSRF) {
	ctx.Data["rulesType"] = model.RuleInfo
	ctx.Data["csrf_token"] = x.GetToken()
	ctx.HTML(200, "addUser")
}

func AddUserDeal(ctx *macaron.Context, ) {
	username := ctx.Req.PostForm.Get("username")
	password := ctx.Req.PostForm.Get("password")
	log.Printf("username:%s;password:%s", username, password)
	err := model.NewUser(username, password)
	if err != nil {
		log.Printf("Fail to insert user{Username:%s,Password:%s} :%s", username, password, err)
		ErrorMsg(ctx, "添加用户失败")
	}
	ctx.Redirect("/index/user/list")
}
