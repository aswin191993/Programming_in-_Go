package main

import "fmt"

func main() {
	defer fmt.Println("world")// deferred call's arguments This can be print in default action at last running wit new line "\n".

	fmt.Println("hello")
	fmt.Println("default Print:")
}

