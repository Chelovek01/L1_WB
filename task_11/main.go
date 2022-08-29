//Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

// Проверка на совпадене значений
func checkExist(arr []string, elem string) bool {
	for _, val := range arr {
		if val == elem {
			return true
		}
	}
	return false
}

func main() {

	//Входные данные
	firstArr := []string{"one", "four", "five", "two", "three", "six", "seven"}
	secondArr := []string{"one", "fifteen", "fourteen", "five", "three"}

	//Результат
	var resultArr []string

	//Проверка на совпадене значений
	for _, val := range firstArr {
		if checkExist(secondArr, val) {
			resultArr = append(resultArr, val)
		}

	}
	fmt.Println(resultArr)

}
