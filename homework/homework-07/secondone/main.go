package main

/*Перепишите программу-конвейер,
ограничив количество передаваемых для обработки значений и обеспечив корректное завершение всех горутин.*/
import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for {
		_, ok := (<-squares)
		if !ok {
			return
		}
		fmt.Println(<-squares)
	}
}
