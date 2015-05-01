package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)//stacking (First in last Out)
	}

	fmt.Println("done")
}

