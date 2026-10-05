// Задание 9: Создай канал done chan
//  struct{} и канал messages chan string.
//  Запусти горутину, которая раз в 500
//  миллисекунд отправляет в messages строку "ping",
//  а через 2 секунды закрывает done.
//  В main в цикле for +
//  select читай messages и печатай их,
//  а при закрытии done выйди из
//  цикла и напечатай "Конец".

package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	messages := make(chan string)

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			messages <- "BOSS"
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		close(done)
	}()

loop:
	for {
		select {
		case v := <-messages:
			fmt.Println("Znacheniya:", v)

		case <-done:
			break loop
		}
	}
	fmt.Println("Kones")
}
