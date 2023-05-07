package main

import "fmt"

func main() {
	C := 0
	fmt.Println("No Of Test Case")
	fmt.Scan(&C)
	for C > 0 {
		x := 0
		fmt.Println("Income")
		fmt.Scan(&x)
		if x > 100 {
			fmt.Println("Income After Tax Deduction = ", x-10)

		} else {
			fmt.Println("Income `Tax Free` = ", x)
		}

		C--
	}

}
