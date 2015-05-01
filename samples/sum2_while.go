package main

import "fmt"

func main(){
	var sum int
	sum = 1
	for sum < 100{
		sum += sum
	}
	fmt.Println("Using While:",sum)
}
