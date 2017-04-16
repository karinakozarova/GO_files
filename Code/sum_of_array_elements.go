package main

import (
	"fmt"
)

func main() {
	var arr [4]int 
	sum := 0
	
	for i := 0; i<4; i++ {
		fmt.Scanf("%d",&arr[i])
	}
	for i := 0; i<4; i++ {
		fmt.Printf("%v",&arr[i])
	}
	for i := 0; i<4; i++ {
		sum += arr[i]
	}
	fmt.Printf("The sum of all array elements is %v",sum)
	
}

