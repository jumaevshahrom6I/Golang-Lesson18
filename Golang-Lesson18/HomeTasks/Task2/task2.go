// Задание 2: Возьми задание 1 и
//  сделай так, чтобы получить оба сообщения:
//  оберни select в цикл for i
//  := 0; i < 2; i++.
//  Напечатай оба сообщения в том порядке,
//  в котором они пришли.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Pervaya"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Vtoraya"
	}()

	for i := 0; i < 2; i++{
		select{
		case v := <-ch1:
			fmt.Println("Value1:", v)

		case v := <-ch2:
			fmt.Println("Value2:", v)
		}
	}
}
