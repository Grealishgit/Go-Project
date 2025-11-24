package main

import (
	"fmt"
)

func loops() {
	for i := 0; i < 100; i += 10 {
		if i == 50 {
			continue
		} else if i == 70 {
			break
		}
		fmt.Println(i)
	}
}

func myMessage() {
	fmt.Println("This is my message function.")
}
