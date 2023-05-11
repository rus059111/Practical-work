package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// создаем каналы для общения между горутинами
	done := make(chan struct{})
	signals := make(chan os.Signal, 1)

	// подписываемся на сигналы прерывания программы
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// горутина вывода квадратов натуральных чисел
	go func() {
		i := 1
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println(i * i)
				i++
				time.Sleep(time.Second)
			}
		}
	}()

	// горутина ожидания сигналов прерывания программы
	go func() {
		<-signals
		close(done)
	}()

	// ожидание завершения вывода квадратов
	<-done
	fmt.Println("выхожу из программы")
}
