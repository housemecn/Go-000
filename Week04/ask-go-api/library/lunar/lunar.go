package lunar

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
)

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: lunar
// @Version: 1.0.0
// @Date: 2020/12/4 00:39
// @Package lunar

// DemoLunar Demo Lunar
func DemoLunar() {
	lunar := calendar.NewLunarFromYmd(1986, 4, 21)
	fmt.Println(lunar.ToFullString())
	fmt.Println(lunar.GetSolar().ToFullString())
}
