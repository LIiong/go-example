package main

import (
	"fmt"
	"strings"
)

//Vertex1 描述
type Vertex1 struct {
	Lat, Long float64
}

var m map[string]Vertex1

var n = map[string]Vertex1{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
/*
	map 
		值为 nil 的 map 是空的，并且不能赋值
		使用make来创建map
*/
func testmap() {
	m = make(map[string]Vertex1)
	m["Bell Labs"] = Vertex1{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	b := make(map[string]int)
	b["a"] = 1
	b["b"] = 4
	fmt.Println("a", b["a"])
	b["a"] = 2
	fmt.Println("a", b["a"])
	//删除元数
	delete(b, "a")
	fmt.Println("a", b["a"])
	//如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
	//同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。
	v,ok := b["b"]
	fmt.Println("v",v,"ok",ok)
}


func wordCount(s string) map[string]int {
	b := make(map[string]int)
	a := strings.Fields(s)
	fmt.Println(a)
	for _, v := range a {
		c := b[v]
		b[v] = c + 1
	}
	return b
}

/*
	Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。 
	函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func testfunc() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
//斐波纳契闭包
func fibonacci() func() int {
	a,b := 0,1
	return func() int {
		t := a
		a, b = b, a+ b
		return t
	}
}

func testfibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}