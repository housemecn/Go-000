package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/guid"
)

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: UnionUID
// @Version: 1.0.0
// @Date: 2020/11/27 21:29
// @Package middleware

// UnionUIDInit 允许接口跨域请求
func UnionUIDInit(r *ghttp.Request) {
	r.SetCtxVar("UnionUID", guid.S())
	r.Middleware.Next()
}
