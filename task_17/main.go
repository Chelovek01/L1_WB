//Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{1, 5, 6, 7, 8, 9, 20, 22, 25}
	x := 5

	i := sort.Search(len(a), func(i int) bool { return x <= a[i] })
	if i < len(a) && a[i] == x {
		fmt.Printf("Найдено %v по индексу %d в %v.\n", x, i, a)
	} else {
		fmt.Printf("Не найдено %v в %v.\n", x, a)
	}
}
