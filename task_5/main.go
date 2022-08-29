package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int)
	var timeOut int
	fmt.Scan(&timeOut)

	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	go sendToChan(arr, ch)
	for range arr {
		time.Sleep(1 * time.Second)
		fmt.Printf("Приняли элемент %v\n", <-ch)

	}

}

func sendToChan(arr []int, ch chan int) {

	for _, val := range arr {
		go func(val int) {
			ch <- val
		}(val)

		fmt.Printf("Отправили %v елемент\n", val)
		time.Sleep(3 * time.Second)

	}

}
