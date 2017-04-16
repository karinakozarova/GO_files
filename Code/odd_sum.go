package main

import "fmt"

func main() {
	var number int 
	var sum = 0
    fmt.Printf("Enter a number:  ")
    fmt.Scanf("%v",&number)
    for i := 1 ; i<=number; i+=2{
    		sum += i
    }
    fmt.Printf("The sum of the odd numbers from 1 to %v is: %v",number,sum)
}
