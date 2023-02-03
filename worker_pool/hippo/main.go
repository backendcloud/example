package main

import "fmt"

func main() {
	tasks := []*Task{
		NewTask(func() (string, error) {
			fmt.Println("Task #1")
			return "Result 1", nil
		}),
		NewTask(func() (string, error) {
			fmt.Println("Task #2")
			return "Result 2", nil
		}),
		NewTask(func() (string, error) {
			fmt.Println("Task #3")
			return "Result 3", nil
		}),
	}

	// hippo.NewWorkersPool(tasks []*Task, concurrency int) *WorkersPool
	p := NewWorkersPool(tasks, 2)
	p.Run()

	var numErrors int
	for _, task := range p.Tasks {
		if task.Err != nil {
			fmt.Println(task.Err)
			numErrors++
		} else {
			fmt.Println(task.Result)
		}
		if numErrors >= 10 {
			fmt.Println("Too many errors.")
			break
		}
	}
}
