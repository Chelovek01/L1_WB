package main

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

import (
	"fmt"
	"strings"
)

func main() {
	//Примеры строк
	str := "test"
	str2 := "abCd"
	str3 := "aabbcc"

	fmt.Println(checkString(str))
	fmt.Println(checkString(str2))
	fmt.Println(checkString(str3))

}

// Функция проверяющая строки
func checkString(str string) bool {
	strings.ToLower(str)
	for i := range str {
		count := strings.Count(str, string(str[i]))
		if count > 1 {
			return false
		}
	}

	return true
}
