package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var count = 10_000_000

func main() {
	//sequence()
	concurrency()
	//concurrency2()
	//concurrency3()
}

func sequence() {

	fmt.Println("sequence")

	start := time.Now()

	func() {
		a := 0
		for i := 0; i < count; i++ {
			a += i
		}

		fmt.Printf("a is %d\n", a)
	}()

	func() {
		b := 0
		for i := 0; i < count; i++ {
			b += i
		}

		fmt.Printf("b is %d\n", b)
	}()

	elapsedTime := time.Since(start)
	fmt.Println("Total Time For Execution: " + elapsedTime.String())
}

func concurrency() {
	fmt.Println("concurrency")

	start := time.Now()

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func(*sync.WaitGroup) {
		defer wg.Done()

		a := 0

		for i := 0; i < count; i++ {
			a += i
		}

		fmt.Printf("a is %d\n", a)
	}(&wg)

	go func(*sync.WaitGroup) {
		defer wg.Done()

		b := 0

		for i := 0; i < count; i++ {
			b += i
		}

		fmt.Printf("b is %d\n", b)
	}(&wg)

	wg.Wait()

	elapsedTime := time.Since(start)
	fmt.Println("Total Time For Execution: " + elapsedTime.String())
}

func concurrency2() {
	fmt.Println("concurrency 2")

	start := time.Now()
	a := 0
	b := 0

	stopA := make(chan bool)
	stopB := make(chan bool)

	go func(a int) {

		for i := 0; i < count; i++ {
			a += i
		}

		fmt.Printf("a is %d\n", a)
		stopA <- true
	}(a)

	go func(b int) {

		for i := 0; i < count; i++ {
			b += i
		}
		fmt.Printf("b is %d\n", b)

		stopB <- true
	}(b)

	<-stopA
	<-stopB

	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution: " + elapsedTime.String())
}

func concurrency3() {
	fmt.Println("concurrency 3")
	runtime.GOMAXPROCS(1)

	start := time.Now()

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()

		a := 0

		for i := 0; i < count; i++ {
			a += i
		}

		fmt.Printf("a is %d\n", a)
	}()

	go func() {
		defer wg.Done()

		b := 0

		for i := 0; i < count; i++ {
			b += i
		}

		fmt.Printf("b is %d\n", b)
	}()

	wg.Wait()

	elapsedTime := time.Since(start)
	fmt.Println("Total Time For Execution: " + elapsedTime.String())
}
