package main

import "fmt"

func main() {
	C := 0
	fmt.Println("No Of Case")
	fmt.Scan(&C)
	for C > 0 {
		x, y := 0, 0
		fmt.Println("Staring Hours")
		fmt.Scan(&x)
		fmt.Println("Ending Hours")
		fmt.Scan(&y)

		if x < y {
			if x >= 1 && x <= 12 {
				if y >= 1 && y <= 12 {
					fmt.Println("Working Hours", y-x)

				} else {
					fmt.Println("Invalid Ending Hours Data")
				}

			} else {
				fmt.Println(" Invalid Staring Hours Data")
			}

		} else {
			fmt.Println("Invalid Hours")
		}
		C--
	}
}
