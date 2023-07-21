package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	//自定义类型和类型别名
	toType()
	//结构体
	toStruct()
	//带标签的结构体
	toStructWithTag()
	//匿名字段和内嵌结构体
	toStructWithAnonymousFields()
	//方法和接收者
	toMethod()

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
	//初始化
	var p person
	p.name = "甘甜"
	p.age = 20
	p.gender = "male"
	p.hobby = []string{"reading", "swimming"}
	fmt.Println(p) // {甘甜 20 male [reading swimming]}
	//简写
	p = person{"qqq", 12, "男", []string{"1", "2"}}
	fmt.Println(p)

	// 用new初始化实例
	n := new(person)
	n.name = "n"
	n.age = 20
	fmt.Println(n) //&{n 20  []}

	//简写 new(Type) 和 &Type{} 是等价的
	n = &person{"www", 12, "男", []string{"111", "2222"}}
	fmt.Println(n)

}

// 带标签的结构体
type desc struct {
	field1 bool   "desc:this is a bool answer"
	field2 string "desc:this is a string answer"
	field3 int    "desc:this is a int answer"
}

func toStructWithTag() {
	t := desc{true, "gold", 10}
	for i := 0; i < 3; i++ {
		showTag(t, i)
	}
}
func showTag(t desc, i int) {
	tType := reflect.TypeOf(t)
	iField := tType.Field(i)
	log.Print(iField)
	log.Print(iField.Tag)
}

type outer struct {
	name string
	age  int
	string
	inner
}
type inner struct {
	region  string
	country string
}

func toStructWithAnonymousFields() {
	person := new(outer)
	person.name = "孙悟空"
	person.age = 100
	person.string = "齐天大圣"
	person.region = "China"
	person.country = "中国"

	log.Print(person)

	//使用结构体字面量
	p2 := outer{"猪八戒", 20, "天蓬元帅", inner{"china", "中国"}}
	log.Print(p2)

	/**
	通过类型 outer.int 的名字来获取存储在匿名字段中的数据，于是可以得出一个结论：
	在一个结构体中对于每一种数据类型只能有一个匿名字段。
	ps:内嵌结构体出现同名时，外层会覆盖内层名字
	**/
}

type animal struct {
	desc string
	name string
}

func (a animal) speak() {
	switch a.desc {
	case "猫":
		log.Print("喵喵喵~")
	case "狗":
		log.Print("汪汪~")
	}
}
func toMethod() {
	/**
	go 方法是作用在接收者上的一个函数，接收者是任何类型的变量（接口除外）
	接收者不能是一个指针类型 但是可以是任何允许类型的指针
	对于一个类型中的方法 不允许重载
	但同名的方法可以出现在多个接收者类型里
	**/
	//实例化
	cat := animal{"猫", "小宝"}
	cat.speak()
}
