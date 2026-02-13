package lesson4

import (
	"fmt"
	"time"
)

func level1_basicParallel() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("A")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("B")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second) 
}

//Wrong
// func level1_loopBug() {
// 	for i := 0; i < 5; i++ {
// 		go func() {
// 			fmt.Println(i)
// 		}()
// 	}
// 	time.Sleep(500 * time.Millisecond)
// }

func level1_loopFix() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(500 * time.Millisecond)
}