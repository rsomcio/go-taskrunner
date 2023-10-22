package task

import "time"

// t := task.New()
// c.Add()
// c.Start() - start first job in queue

type job struct {
	run func()
}

type Task struct {
	ticker *time.Ticker
	jobs   map[string]*job
}

func New() *Task {
	return &Task{
		jobs: map[string]*job{},
	}
}
