package main

import "fmt"

func main() {
	C := 0
	fmt.Println("No Of Test Case")
	fmt.Scan(&C)
	for C>0{
		x,y:=0,0
		fmt.Println("Total Income")
	fmt.Scan(&x)
	fmt.Println("Tax Value")
	fmt.Scan(&y)
	if x>y{
		fmt.Println("Investment Money = ",x-y)
	}else{
		fmt.Println("Invalid Data")
	}

		C--
	}

}