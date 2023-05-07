package main

import "fmt"

func main() {
	C := 0
	fmt.Println("No Of Test Case")
	fmt.Scan(&C)
	for i:=0;i<=C;i++{
		var x int
		var y int
		fmt.Println("Alex Height")
	fmt.Scan(&x)
	fmt.Println("Borden Height")
	fmt.Scan(&y)
	if x>y{
		fmt.Println("Alex Taller Than Borden")

	}else{
		fmt.Println("Borden Taller Than Alex")
	}

	}

}