package main

import "fmt"

func main(){
	var sum int = 1
	for i := 1; i < 10; i++{
		sum += sum
	}
	fmt.Println(sum)
}
