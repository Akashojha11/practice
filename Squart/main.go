package main

import "fmt"

func main() {
	C := 0
	fmt.Println("No Of Test Case")
	fmt.Scan(&C)
	for C>0{
		x:=0
		fmt.Println("No Of Sets")
		fmt.Scan(&x)
		fmt.Println("No Of Squart Complete = ",x*15)

		C--
	}

}