package lesson1

import "fmt"

func Run() {
	var name string = "John"
	fmt.Println("Hello", name)

	name = "Long Nguyen"
	fmt.Println("Hello", name)

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
