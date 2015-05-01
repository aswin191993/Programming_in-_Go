package main

import "fmt"

func main(){
	sum := 0
	for i := 1; i <= 100; i++{//for loop for sum of 1 to 100 numbers
		sum += i
	}
	fmt.Println("sum of 1 to 100:",sum)
}


/*
package main

import "fmt"

func main(){
	sum := 1
	for i := 1; i < 10; i++{
		sum += sum
	}
	fmt.Println(sum)
}
*/
