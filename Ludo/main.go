package main

import "fmt"

func main() {
	C := 0
	fmt.Println("No Of Case")
	fmt.Scan(&C)
	for C > 0 {
		x := 0
		fmt.Println("Number Show On Die")
		fmt.Scan(&x)
		if x == 6 {
			fmt.Println("OPEN")
		} else if x >= 1 && x <= 5 {
			fmt.Println("NO")
		} else if x > 6 {
			fmt.Println("INVALID Number")

		}
		C--
	}

}
