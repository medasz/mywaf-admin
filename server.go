package main

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	"mywaf-admin/router"
	"mywaf-admin/setting"
	"net/http"
)

func init() {
	log.SetPrefix("[mywaf-admin]")
	log.SetFlags(log.Lshortfile)
}

func main() {
	mywaf := macaron.Classic()

	mywaf.Use(macaron.Renderer())
	mywaf.Use(cache.Cacher())
	mywaf.Use(captcha.Captchaer(captcha.Options{
		Width:      200,
		Height:     40,
		Expiration: 60,
	}))
	mywaf.Use(session.Sessioner())
	mywaf.Use(csrf.Csrfer())

	mywaf.Get("/", func(ctx *macaron.Context, sess session.Store) {
		if sess.Get("login") != nil {
			ctx.Redirect("/index/rule/white_ip")
		}
	}, router.LoginHtml)
	mywaf.Post("/login/", csrf.Validate, router.LoginDeal)
	mywaf.Get("/logout/", router.Logout)
	mywaf.Group("/index", func() {
		mywaf.Group("/rule", func() {
			mywaf.Get("/:ruleType", router.GetIndex)
			mywaf.Get("/edit/:Id", router.EditRule)
			mywaf.Post("/edit/:Id", csrf.Validate, router.EditRuleDeal)
			mywaf.Get("/del/:id", router.DeleteRule)
			mywaf.Post("/add/", csrf.Validate, router.AddRuleDeal)
			mywaf.Get("/sync/", router.SyncRule)
		})
		mywaf.Group("/user", func() {
			mywaf.Get("/list/", router.UserList)
			mywaf.Get("/edit/:Id",router.EditUser)
			mywaf.Post("/edit/:Id", csrf.Validate,router.EditUserDeal)
			mywaf.Get("/del/:Id",router.DeleteUser)
			mywaf.Get("/add/",router.AddUser)
			mywaf.Post("/add/", csrf.Validate,router.AddUserDeal)

		})
	}, router.SessionCheck)

	log.Println("mywaf-admin start...")
	log.Println(http.ListenAndServe(setting.WebInfo, mywaf))
}
