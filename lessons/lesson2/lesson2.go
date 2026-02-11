package lesson2

import "fmt"

type User struct {
	ID int
	Name string
	Age int
}

type Student struct {
	User
	Class string
	Score int
}

func Run(i int) {
	var name string = "Long Nguyen"
	fmt.Println("Hello", name, "index", i)
}

func ReplaceSliceIndex() {
	//Add slice 
	a := []int{1, 2, 3}
	fmt.Println("Before replace first index safe", a)
	b := ReplaceFirstIndexSafe(a, 4)
	fmt.Println("After replace first index safe", a)
	fmt.Println("After replace first index safe", b)
}

func ReplaceFirstIndex(slice []int, value int) []int {
	slice[0] = value
	return slice
}

func ReplaceFirstIndexSafe(slice []int, value int) []int {
    copied := make([]int, len(slice))
    copy(copied, slice)

    copied[0] = value
    return copied
}

func PrintUserInfo(user User) {
	fmt.Println("User name:", user.Name)
	fmt.Println("User age:", user.Age)
}

func PrintEventNumber(numbers []int) {
	for _, number := range numbers {
		if number % 2 == 0 {
			fmt.Println(number)
		}
	}
}

func PrintStudentRankAndInfo(student Student) {
	fmt.Println("Student ID:", student.User.ID)
	fmt.Println("Student Name:", student.User.Name)
	fmt.Println("Student Age:", student.User.Age)
	fmt.Println("Student Class:", student.Class)
	fmt.Println("Student Score:", GetStudentRank(student.Score))
}

func GetStudentRank(score int) string {
	switch {
		case score >= 9: 
			return "A"
		case score >= 8:
			return "B"
		default:
			return "c"
	}
}