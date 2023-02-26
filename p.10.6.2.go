package main

//Задание 2.

import (
	"fmt"
)

package main

import (
    "fmt"
    "math"
)

func main() {
    // Запросите у пользователя необходимые данные
    fmt.Print("Введите сумму вклада: ")
    var amount float64
    fmt.Scanln(&amount)

    fmt.Print("Введите ежемесячный процент капитализации: ")
    var rate float64
    fmt.Scanln(&rate)

    fmt.Print("Введите количество лет: ")
    var years int
    fmt.Scanln(&years)

    // Вычислите общее количество месяцев в течение периода вклада
    months := years * 12

    // Вычислите процентную ставку в месяц и переведите ее в доли
    monthlyRate := rate / 100 / 12

    // Вычислите итоговую сумму с учетом начисления процентов
    total := amount * math.Pow(1+monthlyRate, float64(months))

    // Округлите итоговую сумму до целого количества копеек в меньшую сторону
    total = math.Floor(total*100) / 100

    // Вычислите сумму, которая будет зачислена в пользу банка за счет округления копеек
    bankCharges := (total - amount) * 100

    // Выведите результаты на экран
    fmt.Printf("Итоговая сумма: %.2f\n", total)
    fmt.Printf("Сумма, которая будет зачислена в пользу банка: %.0f копеек\n", bankCharges)
}

