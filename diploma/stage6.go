package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

func main() {
	// Чтение файла
	filePath := "Billing.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Удаление пробелов и перевода строки
	mask := strings.TrimSpace(string(content))

	// Преобразование строки в число
	maskInt, err := strconv.ParseUint(mask, 2, 8)
	if err != nil {
		log.Fatal(err)
	}

	// Выделение каждого бита и проверка
	billingData := BillingData{
		CreateCustomer: isBitSet(maskInt, 0),
		Purchase:       isBitSet(maskInt, 1),
		Payout:         isBitSet(maskInt, 2),
		Recurring:      isBitSet(maskInt, 3),
		FraudControl:   isBitSet(maskInt, 4),
		CheckoutPage:   isBitSet(maskInt, 5),
	}

	// Вывод результатов
	fmt.Printf("CreateCustomer: %v\n", billingData.CreateCustomer)
	fmt.Printf("Purchase: %v\n", billingData.Purchase)
	fmt.Printf("Payout: %v\n", billingData.Payout)
	fmt.Printf("Recurring: %v\n", billingData.Recurring)
	fmt.Printf("FraudControl: %v\n", billingData.FraudControl)
	fmt.Printf("CheckoutPage: %v\n", billingData.CheckoutPage)
}

// Проверка бита на установку
func isBitSet(num uint64, bitPos uint) bool {
	mask := uint64(1 << bitPos)
	return (num & mask) != 0
}
