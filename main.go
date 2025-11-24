package main

import (
	"fmt"
)

func main() {
array := []int{1, 2, 3, 4, 5, 6, 7, 8, 0, -1}

	for _, day := range array {
		printDay(day)
	}
	//   loops()  
	//   myMessage()  
	//   myFamilies()
	  structFunc()
}

func printDay(day int) {
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}
