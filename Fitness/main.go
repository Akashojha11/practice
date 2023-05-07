package main

import "fmt"

func main() {
	C := 0

	fmt.Println("No. Of case")
	fmt.Scanf("%d\n", &C)
	for i := 0; i <= C; i++ {
		x ,y:= 0,0
		fmt.Println("Distance Office To Home")
		fmt.Scanf("%d", &x)
		y = x*2 * 5
		fmt.Println("Total Diastance Walking In a Week = ", y)
		

	}

}
