// Задание 13: Напиши программу "проверка сайтов":
//  срез из 5 названий сайтов. Для
//  каждого запусти горутину (через WaitGroup), которая
//  имитирует запрос — спит случайное время
//  от 100 до 1500 миллисекунд и
//  отправляет в канал results строку "<сайт>
//  - OK". В main в цикле
//  for i := 0; i <
//  len(sites); i++ ставь select: либо читаем
//  результат из results, либо срабатывает общий


package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	sites := []string{"google.com", "github.com", "go.dev", "yandex.ru", "habr.com"}
	results := make(chan string, len(sites)) 
	var wg sync.WaitGroup

	for _, s := range sites{
		wg.Add(1)
		go func(site string){
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(1401)+100) * time.Millisecond)
			results <- site + " - OK"
		}(s)
	}

	timeout := time.After(1 * time.Second) 
	count := 0

loop:
	for i := 0; i < len(sites); i++{
		select{
		case v := <-results:
			fmt.Println(v)
			count++
		case <-timeout:
			fmt.Println("Vremya vishlo")
			break loop
		}
	}

	fmt.Println("Otvetili vovremya:", count, "iz", len(sites))
}