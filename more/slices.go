package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}// int = 8bytes
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		pnt:=&p[i]
		fmt.Printf("p[%d] == %d\n address:%p\n", i, p[i],pnt)
	}
}
