package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int ) int {
	sum:=0
	return(func(x int)int{
		sum+=x
		return sum
	})
}

func main() {
	f := fibonacci()
	j:=1
	for i := 0; i < 10; i++ {
		k:=f(j)
		fmt.Println(k)
	}
}

