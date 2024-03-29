Задание 1. Конвейер
Цели задания

   /*  Научиться работать с каналами и горутинами.
    Понять, как должно происходить общение между потоками.

Что нужно сделать

Реализуйте паттерн-конвейер: 

    Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
    Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
    Произведение: следующая горутина умножает квадрат числа на 2.
    При вводе «стоп» выполнение программы останавливается. 

Советы и рекомендации

Воспользуйтесь небуферизированными каналами и waitgroup.
Что оценивается

Ввод : 3

Квадрат : 9

Произведение : 18
Как отправить задание на проверку

Выполните задание в файле вашей среды разработки и пришлите ссылку на архив с вашим проектом через форму ниже.


Задание 2. Graceful shutdown
Цель задания

Научиться правильно останавливать приложения.
Что нужно сделать

В работе часто возникает потребность правильно останавливать приложения. Например, когда наш сервер обслуживает соединения, а нам хочется, чтобы все текущие соединения были обработаны и лишь потом произошло выключение сервиса. Для этого существует паттерн graceful shutdown. 

Напишите приложение, которое выводит квадраты натуральных чисел на экран, а после получения сигнала ^С обрабатывает этот сигнал, пишет «выхожу из программы» и выходит.
Советы и рекомендации

Для реализации данного паттерна воспользуйтесь каналами и оператором select с default-кейсом.
Что оценивается

Код выводит квадраты натуральных чисел на экран, после получения ^С происходит обработка сигнала и выход из программы.
Как отправить задание на проверку

Выполните задание в файле вашей среды разработки и пришлите ссылку на архив с вашим проектом через форму ниже. */


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
