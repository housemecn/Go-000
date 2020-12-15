package middleware

import (
	"net/http"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/guid"

	"git.code.tencent.com/houseme/ask-go-api/library/response"
)

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: error
// @Version: 1.0.0
// @Date: 2020/11/27 21:41
// @Package middleware

// ErrorHandler error handler
func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.WriteJson(response.JSONResponse{
			Code:     500,
			Message:  "哎哟我去，服务器居然开小差了，请稍后再试吧！",
			Data:     nil,
			Time:     gtime.TimestampMicro(),
			UnionUID: r.GetCtxVar("UnionUID", guid.S()).String(),
		})
	}
}
