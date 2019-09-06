package main

import (
	"GoStudy/src/main/basicsStudy/dataType"
	"GoStudy/src/main/basicsStudy/grammar"
	"GoStudy/src/main/basicsStudy/variable"
	"fmt"
)

func main() {
	// 1、基础语法
	grammar.Function()
	// 2、变量声明
	variable.Function()
	// 2、结构体
	man := dataType.GetMan
	fmt.Println(man())
	// 3、匿名函数
	f := func(x, y int) int {
		return x + y
	}
	result := f(1, 2)
	println(result)
}
