package response

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: response
// @Version: 1.0.0
// @Date: 2020/11/27 20:17
// @Package response

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/guid"
)

// JSONResponse 数据返回通用JSON数据结构
type JSONResponse struct {
	Code     int         `json:"code"`     // 错误码((0:成功, 1:失败, >1:错误码))
	Message  string      `json:"message"`  // 提示信息
	Data     interface{} `json:"data"`     // 返回数据(业务接口定义具体数据结构)
	Time     int64       `json:"time"`     // 返回当前响应时间
	UnionUID string      `json:"unionUID"` // 请求的唯一标识
}

// JSON 标准返回结果数据结构封装。
func JSON(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JSONResponse{
		Code:     code,
		Message:  message,
		Data:     responseData,
		Time:     gtime.TimestampMicro(),
		UnionUID: r.GetCtxVar("UnionUID", guid.S()).String(),
	})
}

// JSONExit 返回JSON数据并退出当前HTTP执行函数。
func JSONExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	JSON(r, err, msg, data...)
	r.Exit()
}
