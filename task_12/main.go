//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

//Собственное множетсво любого множества - это такое множество, которое:
//		1) не эквивалентно пустому множеству (где нет элементов);
//		2) не эквивалентно исходному множеству;
//		3) Содержит элемент(ы) из исходного множетсва;

package main

import "fmt"

// Проверка на наличие елмента в результирующем слайсе
func notExist(resultArr []string, element string) bool {

	if len(resultArr) == 0 {
		return true
	}
	for _, val := range resultArr {
		if val == element {
			return false
		}
	}
	return true
}

func main() {
	//Вход данных
	var arr = []string{"cat", "cat", "dog", "cat", "tree"}

	//Результат
	var resultArr []string

	for _, val := range arr {
		//Если элемента нет в resultArr
		if notExist(resultArr, val) {
			resultArr = append(resultArr, val)

		}

	}

	fmt.Println(resultArr)

}
