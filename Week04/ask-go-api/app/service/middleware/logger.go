package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: logger
// @Version: 1.0.0
// @Date: 2020/11/27 21:47
// @Package middleware

// Logger Middleware Log
func Logger(r *ghttp.Request) {
	r.Middleware.Next()
	errStr := ""
	if err := r.GetError(); err != nil {
		errStr = err.Error()
	}
	g.Log().Println(r.Response.Status, r.URL.Path, errStr)
}
