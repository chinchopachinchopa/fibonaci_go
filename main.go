package main

import (
	"fmt"
	"math"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>piska</h1>"))
}

func isPerfectSquare(n int) bool {
	sq := int(math.Sqrt(float64(n)))
	return sq*sq == n
}

func isFibonacci(n int) bool {
	return isPerfectSquare(5*n*n+4) || isPerfectSquare(5*n*n-4)
}

func fibonacci(n int) (int, int) {
	prev := 0
	curr := 1
	for curr < n {
		prev, curr = curr, prev+curr
	}
	return prev, curr
}

func main() {
	port := "8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":"+port, mux)
	var num int
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Ошибка.Введите число.")
		return
	}

	if isFibonacci(num) {
		prev, curr := fibonacci(num)
		fmt.Printf("Числа фибоначи рядом: %d, %d\n", prev, curr+prev)
	} else {
		prev, curr := fibonacci(num)
		if num-prev < curr-num {
			fmt.Printf("Ближайшее число Фибоначчи: %d\n", prev)
		} else {
			fmt.Printf("Ближайшее число Фибоначчи: %d\n", curr)
		}
	}
}
