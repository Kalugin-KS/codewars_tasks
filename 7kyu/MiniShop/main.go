// Описать структуру с именем TOVAR, содержащую следующие поля: NAZV — название товара; STOIM — стоимость товара; KOL — количество товара.

// Написать программу, выполняющую следующие действия:
// • ввод с клавиатуры данных в массив SHOP, состоящий из 5 элементов типа TOVAR;
// • вывод на экран информации о товарах, стоимость которых больше введенного с клавиатуры значения;
// • если таких товаров нет, выдать на дисплей соответствующее сообщение.

package main

import (
	"fmt"
	"strconv"
)

type TOVAR struct {
	NAZV  string
	STOIM int
	KOL   int
}

func main() {

	var SHOP [5]TOVAR
	var nazv, str string
	var stoim, kol, price int

	for i := range SHOP {
		str = strconv.Itoa(i + 1)
		fmt.Print("Заполните поля " + str + " товара: ")
		fmt.Scanf("%s %d %d \n", &nazv, &stoim, &kol)
		SHOP[i] = NewTovar(nazv, stoim, kol)
	}
	fmt.Print("Введите минимальную стоимость: ")
	fmt.Scanf("%d \n", &price)

	haveCheck := false
	for _, val := range SHOP {

		if val.STOIM > price {
			fmt.Printf("Соответствующие товары: %v\n", val)
			haveCheck = true
		}
	}
	if !haveCheck {
		fmt.Println("Таких товаров нет")
	}

}

// Функция конструктор для структуры TOVAR
func NewTovar(nazv string, stoim, kol int) TOVAR {
	return TOVAR{
		NAZV:  nazv,
		STOIM: stoim,
		KOL:   kol,
	}

}
