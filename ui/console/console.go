package console

import (
	"fmt"
	"main/logic"
)

func SMain() {

	var num int
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Ошибка.Введите число.")
		return
	}

	if logic.IsFibonacci(num) {
		prev, curr := logic.Fibonacci(num)
		fmt.Printf("Числа фибоначи рядом: %d, %d\n", prev, curr+prev)
	} else {
		prev, curr := logic.Fibonacci(num)
		if num-prev < curr-num {
			fmt.Printf("Ближайшее число Фибоначчи: %d\n", prev)
		} else {
			fmt.Printf("Ближайшее число Фибоначчи: %d\n", curr)
		}
	}
}
