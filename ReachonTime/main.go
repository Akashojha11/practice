package main

import (
	"fmt"
)

func main() {
	C := 0
	fmt.Println("No Of Case")
	fmt.Scan(&C)
	for C > 0 {
	x := 0
	fmt.Println("Enter Time In Minutes")
	fmt.Scan(&x)
	
		if x >= 30 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		C--
	}
}
