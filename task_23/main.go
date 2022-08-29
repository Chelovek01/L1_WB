package main

import "fmt"

func main() {
	//Пример с набором числе
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	//Переменная для элемента удаления
	var delElem int
	fmt.Printf("Введите индекс элемента для удаления в списке %v. От %v до %v\n", arr, 0, len(arr)-1)
	fmt.Scan(&delElem)

	//Вызов функции удаления элемента
	deleteElement(arr, delElem)

}

// Функция удаления елемента
func deleteElement(arr []int, elem int) []int {
	var result []int

	//Идем по циклу и удаляем елемент
	for i, _ := range arr {
		switch {
		case i == elem && i == 0:
			result = append(arr[1:])
		case i == elem:
			result = append(arr[0:elem], arr[elem+1:]...)
		case i == elem && i == len(arr):
			result = append(arr[0 : elem-1])
		}
	}
	fmt.Println(result)
	return result

}
