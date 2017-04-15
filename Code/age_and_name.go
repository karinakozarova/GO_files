package main

import "fmt"

func main() {
		var age int
		var name string
		fmt.Printf("What's your name? ")
		fmt.Scanf("%v",&name)
		fmt.Printf("What's your age? ")
		fmt.Scanf("%v",&age)
		fmt.Printf("Your age is %v and your name is %v.\n", age, name)
		
}

