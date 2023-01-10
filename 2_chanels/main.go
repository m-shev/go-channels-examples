package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//first()
	//second()
	third()
}

func writer(c chan<- int, start int, end int, delay time.Duration, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	for i := start; i < end; i++ {
		c <- i
		time.Sleep(delay * time.Millisecond)
	}
}

func reader(c <-chan int, name string) {
	for v := range c {
		fmt.Printf("%s got from chanel %d\n", name, v)
	}

	fmt.Println("Reader: finished")
}

func first() {
	c := make(chan int)

	go reader(c, "Reader1")

	writer(c, 0, 10, 500, nil)
}

func second() {
	c := make(chan int)

	go reader(c, "Reader1")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go writer(c, 0, 10, 500, &wg)

	wg.Add(1)
	go writer(c, 10, 20, 250, &wg)

	wg.Wait()
}

func third() {
	c := make(chan int)

	go reader(c, "Reader1")

	go reader(c, "Reader2")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go writer(c, 0, 10, 500, &wg)

	wg.Wait()
}
