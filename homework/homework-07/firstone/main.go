package main

/* Уберите из первого примера (Фибоначчи и спиннер) функцию, вычисляющую числа Фибоначчи.
Как теперь заставить спиннер вращаться в течение некоторого времени? 10 секунд? */
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	command := make(chan string)
	wg.Add(1)
	go spinner(500*time.Millisecond, &wg, command)
	go waitForTime(command)
	wg.Wait()

}
func waitForTime(c chan string) {

	time.Sleep(10 * time.Second)
	c <- "Pause"
}

func spinner(delay time.Duration, wg *sync.WaitGroup, c chan string) {
	defer wg.Done()

	for {
		select {
		default:
			for _, r := range "-\\|/" {
				fmt.Printf("%c\r spinnig for 10 second", r)
				time.Sleep(delay)

			}
		case <-c:
			return
		}

	}
}
