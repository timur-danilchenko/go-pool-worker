package main

import "sync"

type Task func()

type Worker struct {
	ID         int
	BufferSize int
	Tasks      chan Task
}

func (w *Worker) AddTask(m []Task, n int) {
	tasksChan := make(chan Task, len(m))
	wg := sync.WaitGroup{}

	for range n {
		go func() {
			wg.Add(1)
			defer wg.Done()

			for true {
				task, ok := <-tasksChan
				if !ok {
					break
				}
				task()
			}
		}()
	}

	for _, task := range m {
		tasksChan <- task
	}
	close(tasksChan)

	wg.Wait()
}

func main() {
	worker := Worker{
		ID:         1,
		BufferSize: 10,
		Tasks:      make(chan Task),
	}

	go func(worker Worker) {

	}(worker)
}
