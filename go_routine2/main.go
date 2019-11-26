package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go task1(&wg)
	go task2(&wg)
	task3()
	wg.Wait()
}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("Task 1 Done")

}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("Task 2 Done")

}

func task3() {
	time.Sleep(2 * time.Second)
	fmt.Println("Task 3 Done")

}
