package main

type Task func()

func WorkerPool(m []Task, n int) {
	tasksChan := make(chan Task, len(m))
	waitWorkers := make(chan struct{}, n)

	for range n {
		go func() {
			defer func() {
				waitWorkers <- struct{}{}
			}()
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

	for range n {
		<-waitWorkers
	}
}
