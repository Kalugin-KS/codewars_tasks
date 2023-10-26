package calc

import (
	"errors"
)

type calculator struct{}

// Функция-конструктор для структуры calculator
func NewCalculator() calculator {

	return calculator{}

}

// Функция принимает на вход два вещественных числа и знак оператора, возвращает результат и ошибку
func (c *calculator) Calculate(num1, num2 float64, znak string) (float64, error) {

	switch znak {

	case "+":
		return c.plusCalc(num1, num2), nil

	case "-":
		return c.minusCalc(num1, num2), nil

	case "*":
		return c.multiCalc(num1, num2), nil

	case "/":
		if num2 == 0 {
			return 0, errors.New("деление на ноль, введите другие числа")
		}
		return c.divideCalc(num1, num2), nil

	default:
		return 0, errors.New("несуществующий оператор")
	}

}

// Метод для сложения двух чисел
func (c *calculator) plusCalc(num1, num2 float64) float64 {
	return num1 + num2
}

// Метод для вычитания двух чисел
func (c *calculator) minusCalc(num1, num2 float64) float64 {
	return num1 - num2
}

// Метод для перемножения двух чисел
func (c *calculator) multiCalc(num1, num2 float64) float64 {
	return num1 * num2
}

// Метод для деления двух чисел
func (c *calculator) divideCalc(num1, num2 float64) float64 {
	return num1 / num2
}
