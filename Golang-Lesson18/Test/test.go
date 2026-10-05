// package main
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch1 <- "от ch1 (1 сек)"
// 	}()

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch2 <- "от ch1 (1 сек)"
// 	}()

// 	select{
// 	case msg := <- ch1:
// 		fmt.Println("Poluchil:", msg)

// 	case msg := <- ch2:
// 		fmt.Println("Poluchil:", msg)
// 	}
// }



// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch := make(chan string)

// 	go func() {
// 		time.Sleep(3 * time.Second) // долго работает
// 		ch <- "результат"
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Получил", msg)
// 	case <-time.After(2 * time.Second):
// 		fmt.Println("Таймаут: не дождался")
// 	}
// }



// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func work(id int) {
// 	fmt.Println("Работает", id)
// 	time.Sleep(500 * time.Millisecond) // без паузы цикл грузит CPU на 100%
// }

// func main() {
// 	done := make(chan struct{})
// 	var wg sync.WaitGroup

// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			for {
// 				select {
// 				case <-done:
// 					return // канал закрыт, выходим
// 				default:
// 					work(id)
// 				}
// 			}
// 		}(i)
// 	}

// 	time.Sleep(2 * time.Second)
// 	close(done) // сигнал всем горутинам сразу
// 	wg.Wait()   // ждём, пока все выйдут
// 	fmt.Println("Все остановлены")
// }


// package main

// import "fmt"

// func main() {
// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Горутина: работа сделана")
// 		done <- struct{}{} // нажал звонок
// 	}()

// 	<-done // main ждёт звонок
// 	fmt.Println("Main: услышал звонок, иду дальше")
// }



package main

import "fmt"

type Point struct{ X, Y int }

func main() {
	p := Point{1, 2}      // тип Point + значение {1, 2}
	e := struct{}{}       // тип struct{} + значение {}
	fmt.Println(p, e)
}