package main

import "fmt"

func main() {
	T := 0
	fmt.Println("Enter T Value =")
	fmt.Scanf("%d\n",&T)
	for i:= 1; i <= T; i++ {
		x,y,a :=0,0,0
		fmt.Println("Enter x Value =")
	fmt.Scanf("%d\n",&x)
	fmt.Println("Enter y Value =")
	fmt.Scanf("%d\n",&y)
	fmt.Println("Enter a Value =")
	fmt.Scanf("%d\n",&a)
	if x<a && y>=a{
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}
	}
}
