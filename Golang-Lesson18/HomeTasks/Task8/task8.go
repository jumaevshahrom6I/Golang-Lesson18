// Задание 8: Создай два канала chan
//  int: even и odd. Запусти одну
//  горутину, которая в цикле от 1
//  до 10 отправляет чётные числа в
//  even, а нечётные в odd, и
//  после цикла закрывает оба канала (через
//  close). В main прочитай всё через
//  select в цикле for и напечатай
//  "чётное: N" или "нечётное: N". Подсказка:
//  не забудь про v, ok :=

package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		close(even)
		close(odd)
	}()

	for even != nil || odd != nil {
		select {
		case v, ok := <-even:
			if !ok {
				even = nil
				continue
			}
			fmt.Println("Znacheniya juft:", v)

		case v, ok := <-odd:
			if !ok {
				odd = nil
				continue
			}
			fmt.Println("Znacheniya toq:", v)
		}
	}

	fmt.Println("Kones")
}

