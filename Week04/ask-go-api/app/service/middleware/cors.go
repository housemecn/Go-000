package middleware

import (
	"github.com/gogf/gf/net/ghttp"
)

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: cors
// @Version: 1.0.0
// @Date: 2020/11/27 20:53
// @Package middleware

// CORS 允许接口跨域请求
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
