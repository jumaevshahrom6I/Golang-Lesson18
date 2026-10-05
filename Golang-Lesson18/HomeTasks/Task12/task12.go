// Задание 12: Напиши "очередь с отбрасыванием":
//  канал tasks := make(chan int, 3)
//  и горутина-потребитель, которая читает tasks и
//  обрабатывает каждую задачу 300 миллисекунд (time.Sleep).
//  В main в цикле от 1
//  до 10 с паузой 100 миллисекунд
//  отправляй задачи через select с default.
//  Если буфер полон — печатай "Задача
//  отброшена". В конце дождись потребителя через
//  WaitGroup (для этого закрой tasks после


package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tasks := make(chan int, 3)
	var wg sync.WaitGroup
	processed := 0
	dropped := 0

	wg.Add(1)
	go func(){
		defer wg.Done()
		for v := range tasks{
			time.Sleep(300 * time.Millisecond)
			fmt.Println("Obrabotano:", v)
			processed++
		}
	}()

	for i := 1; i <= 10; i++{
		select{
		case tasks <- i:
		default:
			fmt.Println("Zadacha otbroshena:", i)
			dropped++
		}
		time.Sleep(100 * time.Millisecond)
	}

	close(tasks) 
	wg.Wait()
	fmt.Println("Obrabotano:", processed, "Otbrosheno:", dropped)
}