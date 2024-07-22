package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sk string
	fmt.Print("Введите число: ")
	fmt.Scan(&sk)

	ww, err := strconv.Atoi(sk)
	if err != nil {
		fmt.Println("Ошибка: введите корректное число.")
		return
	}

	kk := sumDigits(ww)
	fmt.Printf("Сумма цифр числа %s: %d\n", sk, kk)
}

func sumDigits(fq int) int {
	sum := 0
	for fq > 0 {
		sum += fq % 10
		fq /= 10
	}
	return sum
}
