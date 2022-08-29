// 8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
	"strconv"
)

func changeBit(n int, pos uint) int {
	// Определяем какое значение у бита который нужно сменить
	valueBit := findBit(n, pos)
	if valueBit == 0 {
		// Если 0 используем or mask
		n |= (1 << pos)
	} else {
		// Если 1 используем and not(mask)
		n &= ^(1 << pos)
	}
	return n
}

func findBit(n int, pos uint) (bit int) {
	bit, _ = strconv.Atoi(fmt.Sprintf("%b", n&(1<<pos)))
	return
}

func main() {
	fmt.Println(changeBit(204, 3))
}
