package main

import (
	"fmt"
	"sync"
	"time"
)
/*
Timer 延迟执行
*/
func main20()  {
	time1 := time.NewTimer(2 * time.Second)
	<- time1.C
	fmt.Println("time1 excute")
}
/*
Ticker 循环执行
*/
func main21()  {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <- done :
				return
			case t := <- ticker.C:
				fmt.Println("ticker at ", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stoped")
}

func worker1(i int,a <-chan int, b chan<- int)  {
	for j := range a {
		fmt.Println("worker", i, "start a ", j)
		time.Sleep(time.Second)
		fmt.Println("worker", i, "end a ", j)
		b <- 2 * j
	}
}
/*
线程池，同时启动三个协程
*/
func main22()  {
	const num = 5
	a := make(chan int, num)
	b := make(chan int, num)

	for i := 1; i <= 3; i++ {
		go worker1(i, a, b)
	}
	for i := 1; i <= num; i++ {
		a <- i
	}
	close(a)
	for i := 1; i <= num; i++ {
		<- b
	}
}
func worker2(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker2 is start", i)
	time.Sleep(2 * time.Second)
	fmt.Println("worker2 is finished", i)
}
/*
WaitGroup 等待所有协程执行完毕
*/
func main23()  {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker2(i, &wg)
	}
	wg.Wait()
	fmt.Println("all finished")
}