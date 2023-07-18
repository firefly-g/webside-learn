package main

import (
	"fmt"
)

func main() {
	//自定义类型和类型别名
	toType()
	//结构体
	toStruct()

}

type myInt int     //自定义类型
type yourInt = int //类型别名  为了更清晰知道变量的类型
func toType() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m yourInt
	m = 300
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	//rune为int32的类型别名
	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c) //int32
}

/*
*
使用type和struct关键字来定义结构体
*
*/
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func toStruct() {
	var p person
	p.name = "甘甜"
	p.age = 20
	p.gender = "male"
	p.hobby = []string{"reading", "swimming"}
	fmt.Println(p)

}
