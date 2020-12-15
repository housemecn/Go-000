package router

import (
	"git.code.tencent.com/houseme/ask-go-api/app/api/article"
	"git.code.tencent.com/houseme/ask-go-api/app/api/auth"
	"git.code.tencent.com/houseme/ask-go-api/app/api/hello"
	"git.code.tencent.com/houseme/ask-go-api/app/api/home"
	"git.code.tencent.com/houseme/ask-go-api/app/service/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.UnionUIDInit, middleware.Logger, middleware.CORS, middleware.ErrorHandler)
		group.ALL("/", hello.Hello)
		ctlAuth := new(auth.Controller)
		group.ALL("/auth", ctlAuth)
		ctlHome := new(home.Controller)
		group.ALL("/home", ctlHome)
		ctlArticle := new(article.Controller)
		group.ALL("/article", ctlArticle)

	})

}
