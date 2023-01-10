package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func graceful(start time.Time) {
	fmt.Printf("Good by, last session lasted %.2f seconds\n", time.Since(start).Seconds())
}

func main() {
	//example1()
	//example2()
	//example3()
}

func example1() {
	start := time.Now()

	clear()

	for {
		fmt.Println(time.Now().Local().Format("15:4:5"))
		time.Sleep(1000 * time.Millisecond)
		clear()
	}

	graceful(start)
}

func example2() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	start := time.Now()

	clear()

	for {
		select {
		case <-stop:
			fmt.Println("Got interrupt signal")
			graceful(start)
			return
		default:
			clear()
			fmt.Println(time.Now().Local().Format("15:4:5"))
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func example3() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clear()

Loop:
	for {
		select {
		case <-stop:
			fmt.Println("Got interrupt signal")
			graceful(start)
			break Loop
		case <-ctx.Done():
			fmt.Println("Got cancel signal from context")
			graceful(start)
			break Loop
		default:
			fmt.Println(time.Now().Local().Format("15:4:5"))

			time.Sleep(1000 * time.Millisecond)

			clear()
		}
	}

	fmt.Println("Finished")
}
