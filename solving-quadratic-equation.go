package main

import (
	"fmt"
	"math"
)

// Функции Go способны возвращать сразу несколько значений, что очень удобно для нас.
// Мы будет пользоваться математическими функциями, которые используеют тип float64. Помимо этого, данный тип
// обладающий плавающей точкой, позволит получать более точные значения и решать уравнения с дробными корнями.

// Функция, решающая уравнение будет возвращать количество корней и сами корни. Это нужно, чтобы правильно отличить отсутствие корней, которое оформлено возвратом нуля, и корень, который может быть равен нулю
func main() {
	//Два действительных корня
	fmt.Println(findSolve(3, 7, -6))
	fmt.Println(findSolve(-1, 7, 8))
	//Один действительный корень
	fmt.Println(findSolve(4, 4, 1))
	//Мнимые корни
	fmt.Println(findSolve(2, 1, 1))
	fmt.Println(findSolve(1, -4, 5))

	//Вот пример, как можно использовать эту функцию более грамотно.
	printlnSolve(findSolve(3, 7, -6))
	number, x1, x2, x3, x4 := findSolve(-1, 7, 8)
	printlnSolve(number, x1, x2, x3, x4)
}
func findSolve(a float64, b float64, c float64) (int, float64, float64, float64, float64) {
	//В зависимости от дискриминанта корней может быть разное количество.
	choise := diskriminant(a, b, c)
	if choise > 0 {
		return (two(a, b, choise))
	}
	if choise == 0 {
		return (one(a, b))
	}
	if choise < 0 {
		return (complexKorn(a, b, choise))
	}
	return 0, 0, 0, 0, 0
}

// Функция возвращает дискриминант, ничего необычного.
func diskriminant(a float64, b float64, c float64) float64 {
	return b*b - 4*a*c
}

// Функции one, two созданы для дискриминанта >=0.
func one(a float64, b float64) (int, float64, float64, float64, float64) {
	return 1, (-b) / (2 * a), 0, 0, 0
}
func two(a float64, b float64, d float64) (int, float64, float64, float64, float64) {
	return 2, (-b + (math.Sqrt(float64(d)))) / (2 * a), (-b - (math.Sqrt(float64(d)))) / (2 * a), 0, 0
}

// Если дискриминант<0 уравнение обладает двумя мнимыми корнями, которые легко находятся по формулам.
// x1=(-b+sqrt(D))/(2a)
// x2=(-b-sqrt(D))/(2a)
func complexKorn(a float64, b float64, d float64) (int, float64, float64, float64, float64) {

	return 4, (-b / (2 * a)), ((math.Sqrt(math.Abs(d))) / (2 * a)), (-b / (2 * a)), (((-1) * math.Sqrt(math.Abs(d))) / (2 * a))
}

// Функция печати демонстрирует как следует использовать функцию решения. В зависимости от количества корней, брать столько сколько есть.
func printlnSolve(i int, x1 float64, x2 float64, x3 float64, x4 float64) {
	if i == 1 {
		fmt.Println(x1)
	}
	if i == 2 {
		fmt.Println(x1, x2)
	}
	if i == 4 {
		fmt.Println(x1, x2, x3, x4)
	}
}
