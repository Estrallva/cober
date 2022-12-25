package main

import "fmt"

func main() {
	ch := make(chan string)
	send := "子协程"
	go func() {
		fmt.Println("出现")
		ch <- send
	}()
	receive := <-ch
	fmt.Println(" receive = ", receive)

}
