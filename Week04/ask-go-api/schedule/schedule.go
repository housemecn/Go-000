package schedule

import (
	"fmt"

	"github.com/gogf/gf/frame/g"

	"git.code.tencent.com/houseme/ask-go-api/library/gocolly"
)

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: schedule
// @Version: 1.0.0
// @Date: 2020/12/5 00:01
// @Package schedule

// GoSchedule Go Schedule
func GoSchedule() {
	go gocolly.Colly()
	conn := g.Redis("meta").Conn()
	defer conn.Close()
	go gocolly.SelectCategory(conn)
	go gocolly.SelectArticle(conn)
	gocolly.SelectTheme(conn)
	fmt.Println("搜索完成")
}
