package main

import (
	"fmt"
	"time"
)

type Job struct {
	Fn   func() interface{}
	Name string
}

type WorkPool struct {
	queue chan Job
	size  int
}

func NewPool(poolSize int) *WorkPool {
	queue := make(chan Job)
	return &WorkPool{queue: queue, size: poolSize}
}

func (w *WorkPool) Add(job Job) {
	w.queue <- job
}

func (w *WorkPool) Start() {
	for i := 0; i < w.size; i++ {
		go func() {
			for job := range w.queue {
				fmt.Println("execute job" + " " + job.Name)

				job.Fn()
			}
		}()
	}
}

func main() {
	pool := NewPool(3)
	pool.Start()

	go func() {
		for i := 0; i < 10; i++ {
			job := Job{
				Fn: func() interface{} {
					time.Sleep(1000 * time.Millisecond)
					return struct{}{}
				},
				Name: fmt.Sprintf("Job №%d", i),
			}
			pool.Add(job)
		}
	}()

	time.Sleep(10 * time.Second)
}
