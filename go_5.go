package main
import (
	"fmt"
	"math"
	"os"
	"time"
)
/*
	Go 没有类。然而，仍然可以在任意类型（type）上定义方法。
	方法可以与命名类型或命名类型的指针关联。
*/

//MyFloat 类型
type MyFloat float64
//Abs MyFloat的方法 关联类型
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
//Vertex2 结构体
type Vertex2 struct {
	X, Y float64
}
/*
Abs Vertex2的方法 
	关联指针：有两个原因需要使用指针接收者。
		首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。
		其次，方法可以修改接收者指向的值。
*/
func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//方法的调用
func testfunc1() {
	v := &Vertex2{3, 4}
	fmt.Println(v.Abs())
}
func testfunc2() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
/*
	接口
		接口类型是由一组方法定义的集合。
		接口类型的值可以存放实现这些方法的任何值。
*/

//Abser 接口
type Abser interface {
	Abs() float64
}

func testInterface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex2{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。 由于 Abs 只定义在 *Vertex（指针类型） 上， 所以 Vertex（值类型） 不满足 `Abser`。
	// a = v

	fmt.Println(a.Abs())
}


type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func testInterface1() {
	var w Writer

	// os.Stdout 实现了 Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
//IPAddr 地址
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (v IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", v[0],v[1],v[2],v[3])
}

func testInterface2() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
/*
	错误
		Go 程序使用 error 值来表示错误状态。
		与 fmt.Stringer 类似，`error` 类型是一个内建接口：
		type error interface {
			Error() string
		}
		（与 fmt.Stringer 类似，`fmt` 包在输出时也会试图匹配 `error`。）
		通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 `nil`， 来进行错误处理。
*/

//MyError 自定义异常
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func testError() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
