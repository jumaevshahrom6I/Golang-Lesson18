// Задание 5: Создай буферизированный канал make(chan
//  int, 3). Через select с case
//  ch <- i и default попробуй
//  отправить в него числа от 1
//  до 5 в цикле. Если буфер
//  полон — в default напечатай "Буфер
//  полон, пропускаем число ". Сколько чисел
//  попало в канал, а сколько пропустили?

// package main

// import "fmt"

// func main() {
// 	ch1 := make(chan int, 3)

// 	for i := 1; i <= 5; i++{
// 		select{
// 		case ch1 <- i:
// 			fmt.Println("Znacheniya:", i)

// 		default:
// 			fmt.Println("Bolshe netu mesto!")
// 		}
// 	}
// }



package main
import "fmt"
func main() {
	box := make(chan int, 3)
	for i := 1; i <= 5; i++{
		select{
		case box <- i:
			fmt.Println("Znacheniya:", i)

		default:
			fmt.Println("Bolshe netu mesto!")
		}
	}
}