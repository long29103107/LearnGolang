package main

import (
	"fmt"
	"learn-golang/lessons/lesson2"
)

func main() {
	// lesson1.Run()
	// lesson2.Run(i)
	// lesson2.ReplaceSliceIndex()

	// users := []lesson2.User{
	// 	{Name: "Long Nguyen", Age: 20, Email: "longnguyen@gmail.com"},
	// 	{Name: "Long Nguyen", Age: 20, Email: "longnguyen@gmail.com"},
	// }

	// for index := 0; index < len(users); index++ {
	// 	lesson2.PrintUserInfo(users[index])
	// 	fmt.Println("--------------------------------")
	// }

	// eventNumbers := []int{}
	// for i := 0; i < 30; i++ {
	// 	eventNumbers = append(eventNumbers, i+1)
	// }
	// lesson2.PrintEventNumber(eventNumbers)

	student := []lesson2.Student{
		{
			User:  lesson2.User{ID: 1, Name: "Long Nguyen", Age: 20},
			Class: "A",
			Score: 9,
		},
		{
			User:  lesson2.User{ID: 2, Name: "Toan Thai", Age: 30},
			Class: "B",
			Score: 6,
		},
		{
			User:  lesson2.User{ID: 3, Name: "Tu Nguyen", Age: 20},
			Class: "C",
			Score: 8,
		},
	}

	for _, student := range student {
		lesson2.PrintStudentRankAndInfo(student)
		fmt.Println("--------------------------------")
	}
}
