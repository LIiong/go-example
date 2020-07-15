package main //包名

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
	"runtime"
) //导入包
/*Go中基本类型*/
var (
	/* int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr*/
	c int = 11
	//bool
	d bool = true
	//string
	e      string     = "你好"
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//PI 常量
const PI = 3.14

/*
* swap 方法名
* x, y string 参数 类型在后 = x string, y string
* (string, string) 返回值 支持多返回
* 在 Go 中，首字母大写的名称是被导出的（类似public）。
Foo 和 FOO 都是被导出的名称。名称 foo 是不会被导出的（类似private）。
*/
func swap(x, y string) (string, string) {
	return y, x
}

func deferTime() {
	/*变量
	 * 以下类似于 var a,b string
	 * 初始化：
	 * 1. var i, j int = 1, 2
	 * 2. var c, python, java = true, false, "no!"
	 * 3. c, python, java := true, false, "no!"（在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用 于替代 var 定义。函数外的每个语句都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。）
	 * 变量在定义时没有明确的初始化时会赋值为_零值_。数值类型为 `0`，布尔类型为 `false`，字符串为 `""`（空字符串）
	 */
	//延迟调用 defer
	defer fmt.Println("end")
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
func changeType(x int) string {

	return strconv.Itoa(x)
}
func iterator(start, end int) int {
	sum := 0
	//普通for循环
	for i := start; i < end; i++ {
		sum += i
	}
	//可以省略刁前后的条件
	for sum < 10 {
		sum += sum
	}
	/*死循环
	for {
	}
	*/
	return sum
}

func condition(x, n float64) float64 {
	
	a := math.Pow(x, n)
	//条件判断
	if a < 20 {
		return a
	}
	//带条件条件判断
	if b := a + 10; b > 20 {
		return b
	} else {
		fmt.Printf("this is %f", b)
	}
	return 0
}
func switchTest(){
	//除非以 fallthrough 语句结束，否则分支会自动终止。
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
