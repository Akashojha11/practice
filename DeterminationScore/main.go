package main

import "fmt"

func main() {
	C := 0
	fmt.Println("No. Of Case")
	fmt.Scanf("%d\n", &C)
	for C > 0 {
		F := 10
		x, T, S, n := 0, 0, 0, 0
		fmt.Println("Point")
		fmt.Scanf("%d\n", &x)
		T = x / F
		fmt.Println("Test Case")
		fmt.Scanf("%d\n", &n)

		S = T * n
		fmt.Println("Score Number = ", S)
		C--
	}

}
