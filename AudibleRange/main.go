package main

import "fmt"

func main() {
	T := 0

	fmt.Println("Case =")
	fmt.Scanf("%d\n", &T)
	for T>0 {
		x := 0
		fmt.Println("Hertz")
		fmt.Scanf("%d\n", &x)
		if x >= 67 && x <= 450000 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		T--
	}
	

}
