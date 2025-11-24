package main

import (
	"fmt"
)

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func myFamilies() {
	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")
}
