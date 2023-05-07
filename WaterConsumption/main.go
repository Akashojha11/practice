package main

import "fmt"

func main() {
	C := 0
	fmt.Println("No Of Test Case")
	fmt.Scan(&C)
	for i:=0;i<=C;i++ {
		x := 0
		fmt.Println("Water Consume In a Day")
		fmt.Scan(&x)
		if x >= 3000 {
			fmt.Println("YES")

		} else {
			fmt.Println("NO")
		}

		
	}

}

