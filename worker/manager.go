package worker

import (
	"fmt"
	"sync"

	"github.com/19chonm/461_1_23/fileio"
)

const numworkers = 5 // Total number of workers/goroutines to use

func StartWorkers(urlch <-chan string, worker_output_ch chan<- fileio.WorkerOutput) {
	var wg sync.WaitGroup

	wg.Add(numworkers) // Keep track of the number of goroutines being created
	for i := 0; i < numworkers; i++ {
		// Start each worker
		go func() {
			fmt.Println("worker: Start worker")
			for {
				url, ok := <-urlch
				if !ok { // Channel has been closed
					fmt.Println("worker: Close worker")
					wg.Done()
					return
				}
				runTask(url, worker_output_ch)
			}
		}()
	}

	wg.Wait() // Wait for the threads to finish
	close(worker_output_ch)
}
