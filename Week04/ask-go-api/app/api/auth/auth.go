package auth

import (
	"github.com/gogf/gf/net/ghttp"

	"git.code.tencent.com/houseme/ask-go-api/app/service/auth"
	"git.code.tencent.com/houseme/ask-go-api/library/response"
)

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: auth
// @Version: 1.0.0
// @Date: 2020/11/27 20:10
// @Package auth

// LoginInput  login input
type LoginInput struct {
	AuthAccount string `v:"required#登陆用户名标识不能为空"`
	AuthPass    string `v:"required#登陆密码不能为空"`
}

// Controller auth Controller
type Controller struct {
}

// Authorization user Authorization
func (c *Controller) Authorization(r *ghttp.Request) {
	var data *LoginInput
	if err := r.Parse(&data); err != nil {
		response.JSONExit(r, 10401, err.Error())
	}
	//用户授权
	if resp, err := auth.Auth(r.GetCtx(), data.AuthAccount, data.AuthPass, r.GetClientIp()); err != nil {
		response.JSONExit(r, resp.Code, err.Error(), resp.Data)
	} else {
		response.JSONExit(r, resp.Code, resp.Message, resp.Data)
	}
}
