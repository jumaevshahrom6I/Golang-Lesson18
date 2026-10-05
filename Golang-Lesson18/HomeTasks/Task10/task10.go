// Задание 10: Создай канал results chan
//  int. Запусти 3 горутины (через WaitGroup),
//  каждая отправляет в results своё число
//  (1, 2, 3) после паузы. В
//  main в отдельной горутине дождись wg.Wait()
//  и закрой results. Основной поток читает
//  results через for v := range
//  results и печатает сумму всех чисел.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	result := make(chan int)
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++{
		wg.Add(1)
		go func(n int){
			defer wg.Done()
			time.Sleep(time.Duration(n * 100) * time.Millisecond)
			result <- n
		}(i)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	sum := 0
	for v := range result{
		sum += v
	}
	fmt.Println("Sum:", sum)
}