// Задание 1: Создай два канала chan
//  string: ch1 и ch2. Запусти две
//  горутины: первая через 1 секунду отправляет
//  в ch1 "из первого", вторая через
//  2 секунды отправляет в ch2 "из
//  второго". В main сделай один select
//  и напечатай то, что пришло первым.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "из первого"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "iz vtorogo"
	}()

	select {
	case v := <-ch1:
		fmt.Println("Value1:", v)

	case v := <-ch2:
		fmt.Println("Value2:", v)
	}
}
