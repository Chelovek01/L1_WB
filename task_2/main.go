package main

import (
	"fmt"
	"os"
)

func main() {

	arr := [5]int{2, 4, 6, 8, 10}
	chanel := make(chan int)
	go sqr(arr, chanel)
	for range arr {
		fmt.Println(os.Stdout, <-chanel)
	}

}

func sqr(arr [5]int, chanel chan int) {
	for _, val := range arr {

		go func(val int) {
			result := val * val
			chanel <- result
		}(val)
	}
}
