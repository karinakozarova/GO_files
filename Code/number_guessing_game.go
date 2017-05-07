package main

import (
	"fmt"
    "math/rand"
    "time"
)

func main () {
	var number int //declaration
	var tries int = 0
	random_number := random(1, 9) //gets random number
    for ; ;tries++ {
	    	    //fmt.Printf("Enter your prediction: ")
	    	    fmt.Scanf("%v",&number)
		    	if number == random_number {
		    			break;
		    	} else if number<random_number{
		    				fmt.Printf("Smaller than random num\n")
		    	} else {
		    				fmt.Printf("Bigger than random num\n")
		    	}
		 }
    fmt.Printf("You guessed correctly in %v tries",tries)

}

 func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}
    
