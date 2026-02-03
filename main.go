package main

import (
	// "learn-golang/lessons/lesson1"
	"fmt"
	"learn-golang/lessons/lesson2"
)

func main() {
	// lesson1.Run()
	var a [3]int = [3]int{1, 2, 3}
	for i := 0; i < len(a); i++ {
		lesson2.Run(i)
	}
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
