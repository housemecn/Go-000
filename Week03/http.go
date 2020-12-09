package main

import (
	"fmt"

	"github.com/gogf/gf/frame/g"
)

// @Project: Go-000
// @Author: houseme
// @Description:
// @File: http
// @Version: 1.0.0
// @Date: 2020/12/9 21:33
// @Package main

func main() {
	r, err := g.Client().Post(
		"http://127.0.0.1:8199/form",
		g.Map{
			"submit":   "1",
			"callback": "http://127.0.0.1/callback?url=http://baidu.com",
		},
	)
	if err != nil {
		panic(err)
	} else {
		defer r.Close()
		fmt.Println(r.ReadAllString())
	}
	g.Server().Run()
}
