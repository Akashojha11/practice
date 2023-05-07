package main

import "fmt"

func main() {
	C := 0
	fmt.Println("No Of Test Case")
	fmt.Scan(&C)
	for C > 0 {
		x, y := 0, 0
		fmt.Println("Target Runs")
		fmt.Scan(&x)
		fmt.Println("Runs Made By Team ")
		fmt.Scan(&y)
		fmt.Println("Runs Need To Win Match", x-y)

		C--
	}

}
