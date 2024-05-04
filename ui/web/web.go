package web

import (
	"fmt"
	"main/logic"
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		a := "Введите число:"
		expression := fmt.Sprintf(a)
		fmt.Fprintf(w, "<h1>%s</h1>", expression)
		fmt.Fprintf(w, "<form method='post'><input type='text' name='answer'><input type='submit'></form>")
	} else if r.Method == "POST" {
		r.ParseForm()
		answerStr := r.Form.Get("answer")
		num, err := strconv.Atoi(answerStr)
		if err != nil {
			http.Error(w, "Invalid answer", http.StatusBadRequest)
			return
		}

		if logic.IsFibonacci(num) {
			prev, curr := logic.Fibonacci(num)
			fmt.Fprintf(w, "Числа фибоначи рядом: %d, %d\n", prev, curr+prev)
		} else {
			prev, curr := logic.Fibonacci(num)
			if num-prev < curr-num {
				fmt.Fprintf(w, "Ближайшее число Фибоначчи: %d\n", prev)
			} else {
				fmt.Fprintf(w, "Ближайшее число Фибоначчи: %d\n", curr)

			}
		}
	}
}
func Main() {
	port := "8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":"+port, mux)
}
