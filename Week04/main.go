package main

import (
	"fmt"
)

// @Project: Go-000
// @Author: houseme
// @Description:
// @File: main
// @Version: 1.0.0
// @Date: 2020/12/23 14:51
// @Package Week04

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Println(m)
}
