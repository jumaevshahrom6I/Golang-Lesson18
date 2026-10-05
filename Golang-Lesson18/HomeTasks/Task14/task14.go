// Задание 14: Напиши функцию merge(a, b
//  chan int) chan int, которая запускает
//  горутину и объединяет два входных канала
//  в один выходной: в цикле for
//  + select читает из a и
//  b и отправляет значения в выходной
//  канал. Когда оба входных канала закрыты,
//  закрывает выходной. Чтобы не крутиться вхолостую
//  на закрытом канале, после обнаружения закрытия
//  делай локальную переменную-канал равной nil (a


package main

import "fmt"

func merge(a, b chan int) chan int {
	out := make(chan int)
	go func(){
		defer close(out)
		for a != nil || b != nil{
			select{
			case v, ok := <-a:
				if !ok{
					a = nil
					continue
				}
				out <- v
			case v, ok := <-b:
				if !ok{
					b = nil
					continue
				}
				out <- v
			}
		}
	}()
	return out
}

func gen(nums ...int) chan int {
	ch := make(chan int)
	go func(){
		defer close(ch)
		for _, v := range nums{
			ch <- v
		}
	}()
	return ch
}

func main() {
	for v := range merge(gen(1, 2, 3), gen(10, 20, 30)){
		fmt.Print(v, " ")
	}
	fmt.Println()
}