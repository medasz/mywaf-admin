package main

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	"mywaf-admin/setting"
	"net/http"
)

func init() {
	log.SetPrefix("[mywaf-admin]")
}

func main() {
	mywaf := macaron.Classic()

	mywaf.Use(macaron.Renderer())
	mywaf.Use(cache.Cacher())
	mywaf.Use(captcha.Captchaer(captcha.Options{
		ChallengeNums:    0,
		Width:            0,
		Height:           0,
		Expiration:       0,
		CachePrefix:      "",
		ColorPalette:     nil,
	}))
	mywaf.Use(session.Sessioner())
	mywaf.Use(csrf.Csrfer())

	log.Println("mywaf-admin start...")
	log.Println(http.ListenAndServe(setting.WebInfo, mywaf))
}
