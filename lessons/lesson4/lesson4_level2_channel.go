package lesson4

import (
	"fmt"
	"time"
)

func Level2_channelBasic() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Press Enter to exit...")
}


func Level2_fanOut() {
	ch := make(chan int)

	for w := 1; w <= 3; w++ {
		go func(worker int) {
			for job := range ch {
				fmt.Printf("Worker %d processed %d\n", worker, job)
			}
		}(w)
	}

	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)

	time.Sleep(500 * time.Millisecond)
}
