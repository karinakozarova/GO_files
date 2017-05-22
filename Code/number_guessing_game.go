package main

import (
	"fmt"
    "math/rand"
    "time"
)

func main () {
	var number int //declaration
	var tries int = 0
	randomNumber := random(1, 9) //gets random number
    for ; ;tries++ {
	    	    fmt.Printf("Enter your prediction: ")
	    	    fmt.Scanf("%v\n",&number)
					if number == randomNumber {
							break;
					} else if number<randomNumber{
								fmt.Printf("Smaller than random num\n")
					} else {
								fmt.Printf("Bigger than random num\n")
					}
		 }
    fmt.Printf("You guessed correctly in %v tries",tries+1)

}

 func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}
    
