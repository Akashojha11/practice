package main

import "fmt"

func main() {
	T := 0
	fmt.Println("Enter case")
	fmt.Scanf("%d\n", &T)
	for i := 0; i <= T; i++ {
		a, b := 0, 0
		fmt.Println("Patty")
		fmt.Scanf("%d\n", &a)
		fmt.Println("Bun")
		fmt.Scanf("%d\n", &b)
		if a < b {
			fmt.Println("No of Burgers Prepare = ", a)
		} else {
			fmt.Println("No of Burgers Prepare = ", b)
		}

	}
}
