package main

import (
	"fmt"
)

func main () {
	var number int //declaration
	fmt.Scanf("%v",&number)
	
	var i int
	var digit int
	var sum = 0
	for i = 0;;i++{
		if number >= 0 {
				digit = number%10
				sum += digit
				number = number/10
		} else {
				break;
		}
	}
	var avg int 
	avg = sum/i
	
	if(avg >= 7){
				fmt.Printf("Heavy")
	}else {
				fmt.Printf("Light")
	}
	

	}
    
