package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main13() {

    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")
}

func main14()  {
    msg := make(chan string, 2)
    go func ()  {
        msg <- "ping"
    }()
    go func ()  {
        msg <- "123"
    }()
    go func ()  {
        msg <- "456"
    }()
    fmt.Println(<- msg)
    fmt.Println(<- msg)
}

func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main15() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main16() {

    done := make(chan bool, 1)
    go worker(done)

    <-done
}

func main17()  {
    a := make(chan string)

    go func ()  {
        time.Sleep(2 * time.Second)
        a <- "1"
    }()
    select {
    case res := <- a:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("1 timeout")
    }

    b := make(chan string)
    go func ()  {
        time.Sleep(2 * time.Second)
        b <- "b send"
    }()
    select {
    case msg := <- b:
        fmt.Println(msg)
    case <- time.After(1 * time.Second):
        fmt.Println("b time out")
    default:
        fmt.Println("no received")
    }
}

func main18()  {
    a := make(chan int, 2)
    b := make(chan bool)
    go func() {
        for  {
            i,j := <- a
            if j{
                fmt.Println("received i", i, j)
            } else {
                fmt.Println("received all i")
                b <- true
                return
            }

        }        
    }()
    for i := 1; i <= 3; i++ {
        a <- i
        fmt.Println("sent i", i)
    }
    close(a)
    fmt.Println("sent all i")
    <- b
}

func main19()  {
    a := make(chan string, 3)
    a <- "one"
    a <- "two"
    close(a)
    // it is error,when code this after close channel
    // a <- "three"
    for tem := range a {
        fmt.Println(tem)
    }
}