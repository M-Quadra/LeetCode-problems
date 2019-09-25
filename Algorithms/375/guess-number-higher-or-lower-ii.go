package main

import (
	"fmt"
)

func getMoneyAmount(n int) int {
	if n < 2 {
		return 0
	}

	payAry := make([][]int, n+5)
	for i := 0; i < n+5; i++ {
		payAry[i] = make([]int, n+5)
	}

	for len := 2; len <= n; len++ {
		for st := 1; st <= n; st++ {
			ed := st + len - 1
			if ed > n {
				continue
			}

			pay := 0x7FFFFF
			for i := st; i <= ed; i++ {
				tmp := i + maxInt(payAry[st][i-st], payAry[i+1][ed-i])
				pay = minInt(pay, tmp)
			}
			payAry[st][len] = pay
		}
	}

	return payAry[1][n]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(getMoneyAmount(i))
	}
	fmt.Println(getMoneyAmount(100))
}
