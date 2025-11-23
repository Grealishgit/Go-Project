package main

import (
	"fmt"
)

func main() {
	// var arr1 = [4]int{1, 2, 3, 4}
	// arr2 := [...]int{4, 5, 6, 7, 8, 9, 10}
	// var cars = [4]string{"Volvo", "BMW", "Ford", "Tesla"}
	// cars[0] = "Toyota"

	// slice
	// myslice1 := []string{"Apple", "Banana", "Cherry", "Mango", "Orange", "Grapes"}
	myslice1 := []int{1, 2, 3, 4}
	myslice2 := []int{5, 10, 15, 20}
	myslice3 := append(myslice1, myslice2...)

	// fmt.Println(arr1[0])
	// fmt.Println(arr2)
	// fmt.Println(cars)
	// fmt.Println(len(cars))

	fmt.Println(myslice1)
	fmt.Println("length of my slice", len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice3)

}
