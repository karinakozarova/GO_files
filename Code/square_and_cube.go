package main

import (
	"fmt"
)

func main () {
	var number int //declaration
	var cube int
	var square int
	
	fmt.Scanf("%v\n",&number)
	
	square = number*number
	cube = square*number
	
	fmt.Printf("Square : %v\nCube: %v",square,cube)
	}
    
