package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	var re string
	go func() {
		ch1 <- "goroutine1"
		close(ch1)
	}()
	re = <-ch1
	fmt.Println(re)
	go func() {
		<-ch1
		ch2 <- "goroutine2"
		close(ch2)
	}()
	re = <-ch2
	fmt.Println(re)
	go func() {
		<-ch2
		ch3 <- "goroutine3"
		close(ch3)
	}()
	re = <-ch3
	fmt.Println(re)
	<-ch3
	time.Sleep(time.Second * 2)
	fmt.Println("successful")
}
