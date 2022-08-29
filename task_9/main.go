package main

import (
	"fmt"
	"os"
)

//Разработать конвейер чисел. Даны два канала:
//в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

// Кладем данные из слайса в первый канал
func first(arr []int) <-chan int {
	//Создаем канал
	chanel := make(chan int)
	go func() {
		for _, val := range arr {
			//Кладем данные
			chanel <- val
		}
		//закрываем канал
		close(chanel)

	}()
	return chanel
}

// Берем данные из первого канала и умножаем на 2
func second(from <-chan int) <-chan int {
	//Создаем канал
	chanel := make(chan int)
	go func() {
		for val := range from {
			//кладем умноженные данные
			chanel <- val * 2

		}
		//закрываем канал
		close(chanel)

	}()
	return chanel

}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	start := first(arr)
	result := second(start)

	for val := range result {
		fmt.Println(os.Stdout, val)
	}

}
