// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"os"
)

func main() {
	//Переменная для результата
	result := 0

	//Создаем канал
	chanel := make(chan int)

	//Слайз для примера
	arr := []int{2, 4, 6, 8, 10}

	//Запуск функции для вычисления
	go sum(arr, chanel)

	for range arr {
		data := 0
		data += <-chanel
		result += data
		fmt.Printf("Данные поступающие из канала %v %v\n", os.Stdout, data)

	}

	fmt.Printf("Сумма квадратов чисел %v", result)

}

// Функция для вычисления
func sum(arr []int, chanel chan int) {
	for _, val := range arr {
		go func(val int) {
			chanel <- val * val
		}(val)
	}

}
