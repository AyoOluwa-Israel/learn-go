package main

import "fmt"

func main(){
	x, y , z := "Israel", 4, 4.22
	// var x string 
	// var y int 
	// var z float64

	fmt.Printf("The value %v has a type %T\n",x, x)
	fmt.Printf("The value %v has a type %T\n",y, y)
	fmt.Printf("The value %v has a type %T\n",z, z)
}