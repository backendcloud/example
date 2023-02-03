package main

import "sync"

// Task struct
type Task struct {
	Err    error
	Result string
	f      func() (string, error)
}

// NewTask initializes a new task based on a given work
func NewTask(f func() (string, error)) *Task {
	return &Task{f: f}
}

// Run runs a Task
func (t *Task) Run(wg *sync.WaitGroup) {
	t.Result, t.Err = t.f()
	wg.Done()
}
