package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println("value of v:",v,"\npointer is transparent:value of v.X",v.X,"\tv.Y:",v.Y)
}

