// Задание 3: Создай канал chan int.
//  Запусти горутину, которая через 3 секунды
//  отправит в него число 42. В
//  main сделай select с двумя case:
//  чтение из канала и <-time.After(1 *
//  time.Second). Напечатай "Дождались: 42" или "Таймаут".
//  Поменяй паузы местами и проверь оба
//  варианта.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 7
	}()

	start := time.Now()

	select{
	case v := <-ch:
		fmt.Println(v, time.Since(start).Seconds())

	case <- time.After(1 * time.Second):
		fmt.Println("Taymaut zhdal sec:", time.Since(start).Seconds())
	}
}
