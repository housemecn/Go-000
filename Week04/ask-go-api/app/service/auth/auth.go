package auth

import (
	"context"
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"

	"git.code.tencent.com/houseme/ask-go-api/app/dao"
	"git.code.tencent.com/houseme/ask-go-api/app/model"
	"git.code.tencent.com/houseme/ask-go-api/library/randstring"
	"git.code.tencent.com/houseme/ask-go-api/library/response"
)

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: auth
// @Version: 1.0.0
// @Date: 2020/11/27 20:51
// @Package auth

// Auth user Authorization
func Auth(ctx context.Context, authAccount, authPass, clientIP string) (*response.Response, error) {
	var resp = response.NewDefaultResponse()
	var err error
	var userAuth *model.UserAuth
	userAuth, err = dao.UserAuth.Where("auth_account=?", authAccount).FindOne()
	if err != nil {
		return nil, err
	}
	if userAuth == nil {
		resp.Code = http.StatusNotFound
		resp.Message = http.StatusText(http.StatusNotFound)
		return resp, nil
	}
	if userAuth.AuthPassword != authPass {
		resp.Code = http.StatusForbidden
		resp.Message = http.StatusText(http.StatusForbidden)
		return resp, nil
	}
	userAuth.LoginIp = clientIP
	userAuth.LoginTime = gtime.Now()
	resp.Data = g.Map{"token": randstring.InitRandString(128)}
	return resp, nil
}
