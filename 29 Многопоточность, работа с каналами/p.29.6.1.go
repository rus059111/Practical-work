package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	// создаем каналы для общения между горутинами
	inCh := make(chan int)
	sqCh := make(chan int)
	resCh := make(chan int)

	// создаем waitgroup для ожидания завершения всех горутин
	wg := &sync.WaitGroup{}

	// первая горутина - чтение чисел из стандартного ввода и отправка их в канал inCh
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			var input string
			fmt.Scanln(&input)
			if input == "стоп" {
				break
			}
			num, err := strconv.Atoi(input)
			if err == nil {
				inCh <- num
			}
		}
		close(inCh)
	}()

	// вторая горутина - вычисление квадрата числа и отправка результата в канал sqCh
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range inCh {
			sqCh <- num * num
		}
		close(sqCh)
	}()

	// третья горутина - умножение квадрата на 2 и отправка результата в канал resCh
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range sqCh {
			resCh <- num * 2
		}
		close(resCh)
	}()

	// горутина вывода результатов
	wg.Add(1)
	go func() {
		defer wg.Done()
		for res := range resCh {
			fmt.Println(res)
		}
	}()

	// ожидание завершения всех горутин
	wg.Wait()
}
