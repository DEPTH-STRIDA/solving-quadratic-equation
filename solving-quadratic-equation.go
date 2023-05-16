package main

import (
	"fmt"
	"math"
)

// Функции Go способны возвращать сразу несколько значений, что очень удобно для нас.
// Мы будет пользоваться математическими функциями, которые используеют тип float64. Помимо этого, данный тип
// обладающий плавающей точкой, позволит получать более точные значения и решать уравнения с дробными корнями.
func main() {
	//Два действительных корня
	findSolve(3, 7, -6)
	findSolve(-1, 7, 8)
	//Один действительный корень
	findSolve(4, 4, 1)
	//Мнимые корни
	findSolve(2, 1, 1)
	findSolve(1, -4, 5)
}
func findSolve(a float64, b float64, c float64) {
	choise := diskriminant(a, b, c)
	if choise > 0 {
		fmt.Println(two(a, b, choise))
	}
	if choise == 0 {
		fmt.Println((one(a, b)))
	}
	if choise < 0 {
		fmt.Println(complexKorn(a, b, choise))
	}
}

// Функция возвращает дискриминант, ничего необычного.
func diskriminant(a float64, b float64, c float64) float64 {
	return b*b - 4*a*c
}

// Функции one, two созданы для дискриминанта >=0.
func one(a float64, b float64) float64 {
	return (-b) / (2 * a)
}
func two(a float64, b float64, d float64) (float64, float64) {
	return (-b + (math.Sqrt(float64(d)))) / (2 * a), (-b - (math.Sqrt(float64(d)))) / (2 * a)
}

// Эта функция позволяет находить мнимые корни. Благо go обладает пакетом для работы с такими числами.
// Если дискриминант<0 уравнение обладает двумя мнимыми корнями, которые легко находятся по формулам.
// x1=(-b+sqrt(D))/(2a)
// x2=(-b-sqrt(D))/(2a)
func complexKorn(a float64, b float64, d float64) (complex128, complex128) {
	return complex(-b/(2*a), (math.Sqrt(math.Abs(d)))/(2*a)), complex(-b/(2*a), ((-1)*math.Sqrt(math.Abs(d)))/(2*a))
}
