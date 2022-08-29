package main

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
//значение которых > 2^20.
import (
	"fmt"
	"math/big"
)

func main() {
	//Используем библиотеку big для выполнения операцияй
	x := big.NewFloat(12341444444444444411111111111122222)
	y := big.NewFloat(183945683945673945683945863568934593)

	var result = new(big.Float)

	//Сумма
	result.Add(x, y)
	fmt.Println(result)

	//Разность
	result.Sub(y, x)
	fmt.Println(result)

	//Умножение
	result.Mul(x, y)
	fmt.Println(result)

	//Деление
	result.Quo(y, x)
	fmt.Println(result)

}
