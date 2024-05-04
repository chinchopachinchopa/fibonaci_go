package logic

import (
	"math"
)

func IsPerfectSquare(n int) bool {
	sq := int(math.Sqrt(float64(n)))
	return sq*sq == n
}
func IsFibonacci(n int) bool {
	return IsPerfectSquare(5*n*n+4) || IsPerfectSquare(5*n*n-4)
}
func Fibonacci(n int) (int, int) {
	prev := 0
	curr := 1
	for curr < n {
		prev, curr = curr, prev+curr
	}
	return prev, curr
}
