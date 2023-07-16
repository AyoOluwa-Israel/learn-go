package main

import "fmt"

func main(){
	var x int8 = 127
	fmt.Printf("%v is a type of %T\n", x, x)

	var y uint8 = 255
	fmt.Printf("%v is a type of %T\n", y, y)
}