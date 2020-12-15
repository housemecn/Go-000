package hello

import (
	"github.com/gogf/gf/net/ghttp"
)

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: hello
// @Version: 1.0.0
// @Date: 2020/11/27 20:38
// @Package hello

// Hello is a demonstration route handler for output "Hello World!".
func Hello(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}
