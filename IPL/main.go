package main

import "fmt"

func main() {
	C := 0
	fmt.Println("No. Of Case")
	fmt.Scanln(&C)
	for C > 0 {
		var n, m int = 0, 0
		fmt.Println("No. of Students")
		fmt.Scanln(&n)
		fmt.Println("No. Of Ticket")
		fmt.Scanln(&m)

		if n > m {
			fmt.Println("Remain Ticket = ", n-m)

		} else {
			fmt.Println(0)
		}
		C--

	}

}
