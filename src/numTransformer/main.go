package main

import "fmt"

/*
Сделать конвейер чисел
	- нужно передать n значение через аргументы вызова (--n 10)
	- нужно от 0 до n числа возвести в квадрат асинхронно.
	- переборы значений не должны блокировать основной процесс
*/

func main() {
	squares := make(chan int)

	// here

	for x := range squares {
		fmt.Println(x)
	}
}
