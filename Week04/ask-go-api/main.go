package main

import (
	"github.com/gogf/gf/frame/g"

	_ "git.code.tencent.com/houseme/ask-go-api/boot"
	_ "git.code.tencent.com/houseme/ask-go-api/router"
	"git.code.tencent.com/houseme/ask-go-api/schedule"
)

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: main
// @Version: 1.0.0
// @Date: 2020/11/28 01:02
// @Package main

func main() {
	go schedule.GoSchedule()
	g.Server().Run()
}
