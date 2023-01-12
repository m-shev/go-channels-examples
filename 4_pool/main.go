package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Fn   func() interface{}
	Name string
}

type WorkPool struct {
	sync.Mutex
	queueCh     chan Job
	finishCh    chan any
	size        int
	isStopped   bool
	finishCount int
}

func NewPool(poolSize int) *WorkPool {
	queue := make(chan Job, poolSize)
	finish := make(chan any)
	return &WorkPool{queueCh: queue, size: poolSize, finishCh: finish}
}

func (w *WorkPool) Add(job Job) {
	if !w.isStopped {
		w.queueCh <- job
	}
}

func (w *WorkPool) Start() {
	for i := 0; i < w.size; i++ {
		go func() {
			for job := range w.queueCh {
				fmt.Println("execute job" + " " + job.Name)

				job.Fn()
			}

			w.Lock()
			defer w.Unlock()

			w.finishCount++

			if w.finishCount == w.size {
				w.finishCh <- true
			}
		}()
	}
}

func (w *WorkPool) Stop() {
	w.isStopped = true
	close(w.queueCh)
}

func (w *WorkPool) Finish() <-chan any {
	return w.finishCh
}

func main() {
	pool := NewPool(2)
	pool.Start()

	go func() {
		for i := 0; i < 20; i++ {
			job := Job{
				Fn: func() interface{} {
					time.Sleep(1000 * time.Millisecond)
					return struct{}{}
				},
				Name: fmt.Sprintf("Job №%d", i),
			}
			pool.Add(job)
		}

		pool.Stop()
	}()

	<-pool.Finish()
}
