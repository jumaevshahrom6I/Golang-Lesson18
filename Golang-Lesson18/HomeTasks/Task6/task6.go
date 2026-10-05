// Задание 6: Создай канал chan string
//  с буфером 1 и положи в
//  него одно сообщение. Напиши функцию tryRead(ch
//  chan string), которая через select с
//  default либо печатает прочитанное сообщение, либо
//  "Пусто". Вызови её три раза подряд
//  и посмотри на вывод.

package main

import "fmt"

func tryRead(ch chan string){
	select{
	case v := <- ch:
		fmt.Println("Znacheniya:", v)

	default:
		fmt.Println("Pusto")
	}
}

func main() {
	ch := make(chan string, 1)
	ch <- "BOSS"
	tryRead(ch)
	tryRead(ch)
	tryRead(ch)
}