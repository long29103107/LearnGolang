package lesson4

import (
	"fmt"
	"sync"
	"time"
)

func levelBonus_limitConcurrency() {
	sem := make(chan struct{}, 5) // max 5 concurrent
	var wg sync.WaitGroup

	for i := 1; i <= 20; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			sem <- struct{}{} // acquire
			fmt.Println("Start", id)
			time.Sleep(500 * time.Millisecond)
			fmt.Println("End", id)
			<-sem // release
		}(i)
	}

	wg.Wait()
}