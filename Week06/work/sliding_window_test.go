package work

import (
	"math/rand"
	"testing"
	"time"
)

// @Project: Go-000
// @Author: houseme
// @Description:
// @File: sliding_window_test
// @Version: 1.0.0
// @Date: 2021/1/10 12:06
// @Package work

// TestSlidingWindow .
func TestSlidingWindow(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	sw := NewSlidingWindow(100)
	var r int
	for i := 0; i < 1000; i++ {
		r = rand.Intn(3)
		if r == 1 {
			sw.AddSuccess()
		} else {
			sw.AddFail()
		}
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	}
	t.Log("1秒的bucket长度", sw.Len())
	for _, item := range sw.Data(3) {
		t.Log(item.success, item.fail)
	}
	t.Log("==============")
	for _, item := range sw.Data(5) {
		t.Log(item.success, item.fail)
	}
}
