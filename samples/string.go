package main

import "fmt"

func main(){
	var s string = "hello world!"
	length := len(s)
	fmt.Printf("length:%d,string:%q\n",length,s)
}
