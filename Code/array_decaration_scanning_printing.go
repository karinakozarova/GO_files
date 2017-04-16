package main

import (
	"fmt"
)

func main() {
	var arr [4]int 
	for i := 0; i<4; i++ {
		fmt.Scanf("%d",&arr[i])
	}
	for i := 0; i<4; i++ {
		fmt.Printf("%v",&arr[i])
	}
}
