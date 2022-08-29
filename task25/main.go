package main

//Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

// Функция sleep
func sleepFunc(t int) {
	<-time.After(time.Duration(t) * time.Second)
	fmt.Println("Sleep Over.....")

}

func main() {

	fmt.Println("Start sleep")
	sleepFunc(5)
	fmt.Println("Choose next operation")

}
