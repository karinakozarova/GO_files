package main

import "fmt"

const num_count =  2

func main() {
		var a float64
		var b float64
		fmt.Printf("Enter num1 ")
		fmt.Scanf("%v",&a)
		fmt.Printf("Enter num2 ")
		fmt.Scanf("%v",&b)
		fmt.Printf("Sum: %v" , sum(a,b))
		fmt.Printf("\nAverage: %v ", average(a,b)) //12
}

func sum(a float64,b float64 ) float64 {
	sum := a
	sum += b
	return sum
}

func average(a float64, b float64) float64 {
	var sum float64 = a + b 
	var average float64 = sum / num_count
	return average
}

