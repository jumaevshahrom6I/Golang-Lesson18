// Задание 11: Напиши функцию fastest(), которая
//  запускает три горутины-"сервера" (каждая спит случайное
//  время от 100 до 1000 миллисекунд
//  через rand.Intn и отправляет в общий
//  канал своё имя "сервер 1", "сервер
//  2", "сервер 3"). В main через
//  select прочитай первый ответ из канала
//  и напечатай, какой сервер ответил быстрее
//  всех. Добавь в select ограничение: если
//  никто не ответил за 500 миллисекунд

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fastest() {
	ch := make(chan string, 3) 

	for i := 1; i <= 3; i++{
		go func(n int){
			time.Sleep(time.Duration(rand.Intn(901)+100) * time.Millisecond)
			ch <- fmt.Sprintf("server %d", n)
		}(i)
	}

	select{
	case v := <-ch:
		fmt.Println("Samiy bistriy:", v)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Vse servera slishkom medlennie")
	}
}

func main() {
	fastest()
}