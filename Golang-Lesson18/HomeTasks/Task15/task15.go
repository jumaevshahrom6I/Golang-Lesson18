// Задание 15: Мини-проект "стоп-кран": запусти 3
//  воркера-горутины через WaitGroup. Каждый воркер в
//  бесконечном цикле for + select: если
//  пришёл сигнал из канала stop chan
//  struct{} — печатает "Воркер остановлен" и
//  выходит (return); иначе (в default) печатает
//  "Воркер работает" и спит 400 миллисекунд.
//  В main подожди 2 секунды, затем
//  закрой stop (close(stop)) и дождись, пока
//  все воркеры завершатся (wg.Wait()). Напечатай "Все


package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, stop chan struct{}, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		select{
		case <-stop:
			fmt.Println("Vorker ostanovlen:", id)
			return
		default:
			fmt.Println("Vorker rabotaet:", id)
			time.Sleep(400 * time.Millisecond)
		}
	}
}

func main() {
	stop := make(chan struct{})
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++{
		wg.Add(1)
		go worker(i, stop, &wg)
	}

	time.Sleep(2 * time.Second)
	close(stop)
	wg.Wait()
	fmt.Println("Vse vorkery ostanovleny")
}