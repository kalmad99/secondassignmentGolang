package main

import (
	"fmt"
	"strconv"
)

func main() {
	fibo(15)
}

func fibo(num int) {
	total := 0
	counter := 1
	for num > 0 {
		counter, total = total, total+counter
		fmt.Print(strconv.Itoa(counter) + " ")
		num--
	}
}
