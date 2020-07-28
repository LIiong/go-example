package main

import "fmt"
/*
	类型 [n]T 是一个有 n 个类型为 T 的值的数组。
	表达式 var a [10]int 定义变量 a 是一个有十个整数的数组。
	数组的长度是其类型的一部分，因此数组不能改变大小
*/
func list() {
	//初始化方式一
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	//初始化方式二
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	//遍历
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
	/*切片特性 （左闭右开）
		slice 可以重新切片，创建一个新的 slice 值指向相同的数组。
		表达式
			s[lo:hi]
		表示从 lo 到 hi-1 的 slice 元素，含两端。
	*/
	fmt.Println("p[1:4] ==", p[1:4])
	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])
	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])
}

func testSlice()  {
	/*
	构造slice 
		slice 的零值是 `nil`。
		一个 nil 的 slice 的长度和容量是 0
	*/
	a := make([]int, 5)
	printSlice("a", a) //a len=5 cap=5 [0 0 0 0 0]
	b := make([]int, 0, 5)
	printSlice("b", b) //b len=0 cap=5 []
	c := b[:2]
	printSlice("c", c) //c len=2 cap=5 [0 0]
	d := c[2:5]
	printSlice("d", d) //d len=3 cap=3 [0 0 0]
	var z []int
	fmt.Println(z, len(z), cap(z)) //nil
	if z == nil {
		fmt.Println("z is nil!")
	}
	//向slice中添加元数
	z = append(z, 0, 1,2,3,4)
	fmt.Println("z", z)
	//for range 可以实现对slice和map的遍历
	for i, v := range z[:3] {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	//通过索引获取
	for i := range z {
		fmt.Printf("%d\n", z[i])
	}
	// for _, value := range z {
	// 	fmt.Printf("%d\n", value)
	// }
	maps := map[string]string{"a":"aa","b":"bb"}
	for k,v := range maps {
		fmt.Printf("%s -> %s\n",k, v)
	}
	// fmt.Println(maps)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printVar(v interface{}) {
    fmt.Println(v)
}

func main11() {
    printVar("hello world")
    printVar(false)
    printVar(20200716)
    printVar([]int{1, 2, 3}) 
    // 上面都是打印很简单的数据类型，我们打印个map试下
    m := make(map[int]string) 
	m[9] = "the 9th day"
}