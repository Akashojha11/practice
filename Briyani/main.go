package main

import "fmt"

func Bri() {
	T := 0
	fmt.Println("Enter No Case")
	fmt.Scanf("%d\n", &T)
	for T > 0 {
		X, Y, ans := 0, 0, 0
		fmt.Println("Week Number")
		fmt.Scanf("%d\n", &X)
		fmt.Println("Coins per Week")
		fmt.Scanf("%d\n", &Y)

		ans = X * Y
		fmt.Println("Total Fees", ans)
		T--

	}
}

func main() {
Bri()
}
