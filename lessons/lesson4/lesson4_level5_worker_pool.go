package lesson4

import (
	"fmt"
	"time"
)

func level5_workerPool() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	worker := func(id int, jobs <-chan int, results chan<- int) {
		for j := range jobs {
			time.Sleep(200 * time.Millisecond)
			results <- j * 2
		}
	}

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= 10; i++ {
		fmt.Println("Result:", <-results)
	}
}