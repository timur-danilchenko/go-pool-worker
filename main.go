package main

import "sync"

type Task func()

func WorkerPool(m []Task, n int) {
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
