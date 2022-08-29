package main

import (
	"fmt"
	"os"
)

func main() {
	result := 0
	chanel := make(chan int)
	arr := []int{2, 4, 6, 8, 10}
	go sum(arr, chanel)
	for range arr {
		data := 0
		data += <-chanel
		result += data
		fmt.Printf("Данные поступающие из канала %v %v\n", os.Stdout, data)

	}

	fmt.Printf("Сумма квадратов чисел %v", result)

}

func sum(arr []int, chanel chan int) {
	for _, val := range arr {
		go func(val int) {
			chanel <- val * val
		}(val)
	}

}
