package main

import (
	"fmt"
	"time"
)

func main() {
	//example1()
	//example2()
	//example3()
	//example4()
	example5()
}

func example1() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

func example2() {
	c := make(chan int)

	go func() {
		c <- 1
	}()

	fmt.Println(<-c)
}

func example3() {
	c := make(chan int, 1)
	c <- 1
	fmt.Println(<-c)
}

func example4() {
	c := make(chan int, 3)
	writingFinished := make(chan interface{})
	readingFinished := make(chan interface{})

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Printf("\nSent to the channel %d\n", +i)
		}
		fmt.Printf("\nFinished writing to the channel\n")
		writingFinished <- true
	}()

	go func() {
		for v := range c {
			fmt.Printf("\nRead from the channel %d\n", v)
			time.Sleep(1000 * time.Millisecond)
		}
		fmt.Printf("\nAll data from the channel has been read\n")
		readingFinished <- true
	}()

	<-writingFinished

	fmt.Printf("\nClose the channel\n")

	// Попытка записи в закрытый канал приведёт к панике.
	// Читать из закрытого канала можно.
	close(c)

	<-readingFinished
}

func example5() {
	c := make(chan string)

	go func() {
		c <- "some text"
		close(c)
	}()

	v, isOpen := <-c
	fmt.Printf("Value from the channel \"%v\", is open indicator: %v\n", v, isOpen)

	v, isOpen = <-c
	fmt.Printf("Value from the channel \"%v\", is open indicator: %v\n", v, isOpen)
}
