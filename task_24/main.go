package main

//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
//с инкапсулированными параметрами x,y и конструктором

import (
	"fmt"
	"math"
)

// Структура точки

type Point struct {
	x float64
	y float64
}

// Конструктор
func newPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func main() {
	//Задаем точки
	p1 := newPoint(1, 2)
	p2 := newPoint(5, 7)
	fmt.Printf("Расстояние между точек %.2f", distance(p1, p2))

}

// Функция для расчета расстояния
func distance(p1, p2 Point) float64 {
	result := math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
	return result
}
