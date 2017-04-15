package main
import "fmt"
func main() {
		var a int
		var b int
		fmt.Printf("Enter num1 ")
		fmt.Scanf("%v",&a)
		fmt.Printf("Enter num2 ")
		fmt.Scanf("%v",&b)
		fmt.Printf("Sum: %v" , sum(a,b))
}

func sum(a int,b int ) int {
	sum := a
	sum += b
	return sum

}

