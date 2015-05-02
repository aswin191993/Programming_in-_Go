package main

import "fmt"

func celsius(x float32) float32{
	c := ((x - 32) * 5/9 ) 
	return c
}

func main(){
	fmt.Printf("Fara:122 , Celsius:%v\n", celsius(122))
}	
