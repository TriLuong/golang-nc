package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go task1(messages)
	go task2(messages)
	go task3(messages)
	msg1 := <-messages
	msg2 := <-messages
	msg3 := <-messages
	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)

}

func task1(ch chan string) {
	time.Sleep(2 * time.Second)

	ch <- "task 1 done!"
}

func task2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "task 2 done!"
}

func task3(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "task 3 done!"

}
