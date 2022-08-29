//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"os"
)

func main() {

	//Слайс для примера
	arr := []int{2, 4, 6, 8, 10}
	chanel := make(chan int)

	//Конкурентные вичисления
	go sqr(arr, chanel)
	for range arr {
		fmt.Println(os.Stdout, <-chanel)
	}

}

// Функция для записи в канал вычислений
func sqr(arr []int, chanel chan int) {
	for _, val := range arr {

		go func(val int) {
			result := val * val
			chanel <- result
		}(val)
	}
}
