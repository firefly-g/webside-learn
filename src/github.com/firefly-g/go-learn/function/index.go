package main

import (
	"fmt"
)

func main() {
	//使用defer延迟运行语句 多用于在函数结束之前释放资源*（文件句柄、数据库连接、socket连接）
	// toDeferTest()
	//经典应用
	// toDefetFunc()
	//作用域
	// toContext()
}
func toDeferTest() {
	fmt.Println("start")
	defer fmt.Println("我是一")
	defer fmt.Println("我是二")
	defer fmt.Println("我是三")
	fmt.Println("end")
}
func toDefetFunc() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
} //5
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
} //6

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //c此时等于 y = x = 5
} //5
func f4() (x int) {
	defer func(x int) {
		x++
	}(x) //将x传入 此时为值传递 不会改变原来的变量
	return 5
} //5

var x = 100

func toContext() {
	print1()
}
func print1() {
	x := 200
	fmt.Println(x)
}
