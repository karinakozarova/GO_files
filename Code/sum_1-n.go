package main

import "fmt"

func main() {
	var number int 
	var sum = 0
    fmt.Printf("Enter a number:  ")
    fmt.Scanf("%v",&number)
    for i := 1 ; i<=number; i++{
    		sum += i
    }
    fmt.Printf("The sum of the numbers from 1 to your number is: %v",sum)
}
