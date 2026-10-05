// Задание 4: Создай пустой канал chan
//  int (в него никто ничего не
//  пишет). Сделай select с case <-ch
//  и default, в default напечатай "Данных
//  пока нет". Убедись, что программа не
//  зависает.

package main

import "fmt"

func main() {
	msg := make(chan int)

	select {
	case v := <-msg:
		fmt.Println(v)
	default:
		fmt.Println("Soobsheniya netu")
	}
}