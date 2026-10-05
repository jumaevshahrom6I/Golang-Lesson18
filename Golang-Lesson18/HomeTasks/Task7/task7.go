// Задание 7: Напиши функцию worker(id int,
//  wg *sync.WaitGroup), которая через defer wg.Done()
//  спит id * 100 миллисекунд и
//  печатает "Воркер закончил". Запусти 5 воркеров
//  в цикле и дождись всех через
//  sync.WaitGroup (Add до go, Wait в
//  конце main).

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func worker(id int, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	time.Sleep(time.Duration(id * 200) * time.Millisecond)
// 	fmt.Printf("Rabotnik: %d - gotov\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 1; i <= 5; i++{
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("Gotova")
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(time.Duration(id * 100) * time.Millisecond)
	fmt.Printf("Rabotnik: %d gotov\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++{
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Gotova")
}