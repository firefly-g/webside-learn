package main

import (
	"fmt"
	"log"
)

/*
go中的接口非常灵活 它可以实现很多面向对象的特性，它提供了一种方式来说明对象的行为
*/
type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}
type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}
func (s *Square) Area() float32 {
	return s.side * s.side
}
func main() {
	//sq1 := Square{4}
	sq1 := new(Square)
	sq1.side = 5
	var shaperInt Shaper
	shaperInt = sq1
	fmt.Printf("The square has area: %f\n", shaperInt.Area())
	//接口实现多态
	toInterfacesPoly()
	//接口嵌套接口
	toNestedInterface()
	//类型断言：如何检测和转换接口变量的类型
	toTransTypeOfInterface()
	//类型判断
	toTypeSwitch()

}

func toInterfacesPoly() {
	s := &Square{5}
	q := Rectangle{10, 3}
	//将结构变量存入数组中 赋值给接口变量
	shapes := []Shaper{s, q}
	//遍历数组 一次执行方法
	for i, _ := range shapes {
		log.Print("shape detail:", shapes[i])
		log.Print("Area of this shape is:", shapes[i].Area())
	}

}

type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}
type Lock interface {
	Lock()
	UnLock()
}
type File interface {
	ReadWrite
	Lock
	Close() string
}
type Buffer struct {
	name string
	size float32
}
type Byte struct {
	name string
	size float32
}

func (b Buffer) Read(br Buffer) bool {
	log.Print("读文件中")
	return true
}
func (b Buffer) Write(br Buffer) bool {
	log.Print("取文件中")
	return true
}
func (b Buffer) Close() string {
	log.Print("关闭buffer")
	return "ok"
}
func (b Buffer) Lock() {
	log.Print("上锁")
}
func (b Buffer) UnLock() {
	log.Print("解锁")
}
func (b Byte) UnLock() {
	log.Print("解锁")
}
func toNestedInterface() {
	/***
	一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。
	**/
	f := Buffer{"文件", 1024}
	var fileIntf File
	fileIntf = f
	log.Print(fileIntf.Write(f))

}

func toTransTypeOfInterface() {
	//接口中的类型如果是动态的 可以用类型断言来测试再某个时刻接口是否包含类型T的值
	var b = Buffer{"文件", 1024}
	var fileIntf File
	fileIntf = b
	//判断buffer类型是否存在在接口里
	if t, ok := fileIntf.(Buffer); ok {
		log.Print("the type of areaIntf is ", t)
	}
	//if t, ok := fileIntf.(Byte); ok {
	//	log.Print("areaIntf type include:", t)
	//} else {
	//	log.Print("areaIntf type exclude")
	//}
}

func toTypeSwitch() {
	var b = Buffer{"文件", 1024}
	var fileIntf File
	fileIntf = b
	switch t := fileIntf.(type) {
	case Buffer:
		log.Print("类型是：", t)
	//case Byte:
	//	log.Print("类型是：",t)
	case nil:
		log.Print("类型为空")
	default:
		log.Print("未知的类型")
	}
	//测试一个值是否实现了某个接口
	//if sv, ok := b.(File); ok {
	//	log.Print("v包含了File接口中的Close（）", sv.Close())
	//
	//}
}
