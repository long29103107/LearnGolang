package lesson4

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)	

func level3_waitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
			fmt.Println("Done", id)
		}(i)
	}

	wg.Wait()
}