package main

import (
	"sync"
)

// WorkersPool struct
type WorkersPool struct {
	Tasks       []*Task
	concurrency int
	tasksChan   chan *Task
	wg          sync.WaitGroup
}

// NewWorkersPool initializes a new pool with the given tasks
func NewWorkersPool(tasks []*Task, concurrency int) *WorkersPool {
	return &WorkersPool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}
}

// Run runs all work within the pool and blocks until it's finished.
func (w *WorkersPool) Run() {
	for i := 0; i < w.concurrency; i++ {
		go w.work()
	}

	w.wg.Add(len(w.Tasks))
	for _, task := range w.Tasks {
		w.tasksChan <- task
	}

	// all workers return
	close(w.tasksChan)

	w.wg.Wait()
}

// The work loop for any single goroutine.
func (w *WorkersPool) work() {
	for task := range w.tasksChan {
		task.Run(&w.wg)
	}
}
