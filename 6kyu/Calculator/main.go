// Необходимо разработать простой калькулятор с использованием структур и методов.
// Внутри функции main необходимо из консоли считать 2 числа и оператор. Затем создать экземпляр структуры calculator из пакета calc и вызвать метод Calculate,
// передав ему полученные из консоли значения. Полученный из метода Calculate результат нужно распечатать в консоль.

package main

import (
	"calculator/calc"
	"fmt"
)

func main() {

	var num1, num2 float64
	var znak string

	c := calc.NewCalculator() // Создаем экземпляр структуры calculator, чтобы для нее можно было вызвать экспортируемые методы

	fmt.Println("Введите первое число:")
	_, err := fmt.Scanf("%f\n", &num1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Введите второе число:")
	_, err = fmt.Scanf("%f\n", &num2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Введите оператор:")
	_, err = fmt.Scanf("%s\n", &znak)
	if err != nil {
		fmt.Println(err)
	}

	result, err := c.Calculate(num1, num2, znak)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Ответ равен: %0.2f", result)
	}

}
