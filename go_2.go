package main

import "fmt"
/*
	Go 具有指针。 指针保存了变量的内存地址。
	类型 *T 是指向类型 T 的值的指针。其零值是 `nil`
	& 符号会生成一个指向其作用对象的指针
	* 符号表示指针指向的底层的值
*/
func pointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
//Vertex 结构体
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p1  = &Vertex{1, 2} // 类型为 *Vertex
)

func stuctTest()  {
	v := Vertex{1,2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)
	p := &v
	p.X = 1e9
	fmt.Println(v)
	p1.X = 5
	fmt.Println(v1, p1, v2, v3)
}
