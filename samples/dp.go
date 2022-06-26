package main

import "fmt"

var (
	fib1count = 0
	fib2count = 0
)

func fib(n int) int {
	fib1count++
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

//меморизация
func fib2(n int, mem map[int]int) int {
	fib2count++
	if n <= 2 {
		return 1
	}
	if _, ok := mem[n]; ok {
		return mem[n]
	}
	res := fib2(n-1, mem) + fib2(n-2, mem)
	mem[n] = res
	return res

}

//подход динамического программирования
func fib3(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	mem := make(map[int]int)
	fmt.Println(fib(25))
	fmt.Println(fib2(25, mem))
	fmt.Println(fib3(25))
	fmt.Println("fib1=", fib1count, "fib2=", fib2count)
}
