package lesson4

import (
	"fmt"
	"time"
)

func level2_channelBasic() {
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
}


func level2_fanOut() {
	ch := make(chan int)

	for w := 1; w <= 3; w++ {
		go func(worker int) {
			for job := range ch {
				fmt.Printf("Worker %d processed %d\n", worker, job*2)
			}
		}(w)
	}

	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)

	time.Sleep(500 * time.Millisecond)
}
