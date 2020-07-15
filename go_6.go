package main

import (
    "fmt"
    "sort"
)

func main12() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}

func recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursion(n - 1)
}
/*
sort 排序
*/
func sortTest()  {

    a := []string{"a","b","d","c","e","f"}
    sort.Strings(a)
    fmt.Println(a)
    b := []int{3,2,1,4,6,5}
    sort.Ints(b)
    fmt.Println(b)
    c := sort.IntsAreSorted(b)
    fmt.Println(c)
}
/*
自定义排序
*/
type sLen []string

// sLen 实现 sort.Sort的接口 Interface
func (s sLen) Len() int {
    return len(s)
}
func (s sLen) Less(i, j int) bool{
    return len(s[i]) < len(s[j])
}

func (s sLen) Swap(i, j int)  {
    s[i], s[j] = s[j], s[i]
}

func sortTest1()  {
    a := []string{"abc","a","ab","bcde","bacdef"}
    sort.Sort(sLen(a))
    fmt.Println(a)
}