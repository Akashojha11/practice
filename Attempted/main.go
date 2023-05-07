package main

import "fmt"

func main() {
	C := 0
	fmt.Println("No Of Case")
	fmt.Scanln(&C)
	for C > 0 {
		x, y := 0, 0
		fmt.Println("Total Question")
		fmt.Scanln(&x)
		fmt.Println("Attempted Question")
		fmt.Scanln(&y)
		fmt.Println("Unattempted Question = ", x-y)
		C--
	}
}
