package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ATM struct {
	PIN        string
	Balance    float64
	IsLoggedIn bool
}

func readLine(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода.")
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}

func readAmount(reader *bufio.Reader) (float64, error) {
	input := readLine(reader)
	amount, err := strconv.ParseFloat(input, 64)
	if err != nil || amount <= 0 {
		return 0, fmt.Errorf("Неверная сумма")
	}
	return amount, nil
}

func main() {
	atm := &ATM{
		PIN:     "1234",
		Balance: 1000.00,
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите PIN-код: ")
	for attempts := 0; attempts < 3; attempts++ {
		pin := readLine(reader)
		if pin == atm.PIN {
			atm.IsLoggedIn = true
			fmt.Println("Вход выполнен успешно!")
			break
		}
		remaining := 2 - attempts
		fmt.Printf("Неверный PIN-код. Осталось попыток: %d\n", remaining)
		if remaining == 0 {
			fmt.Println("Карта заблокирована. Завершение работы.")
			return
		}
		fmt.Print("Введите PIN-код: ")
	}

	if !atm.IsLoggedIn {
		return
	}

	for {
		fmt.Println("\n=== БАНКОМАТ ===")
		fmt.Println("1. Проверить баланс")
		fmt.Println("2. Пополнить счёт")
		fmt.Println("3. Снять наличные")
		fmt.Println("4. Выйти")
		fmt.Print("Выберите действие: ")

		choice := readLine(reader)

		switch choice {
		case "1":
			fmt.Printf("Ваш баланс: %.2f руб.\n", atm.Balance)

		case "2":
			fmt.Print("Введите сумму для пополнения: ")
			amount, err := readAmount(reader)
			if err != nil {
				fmt.Println("Введите корректную положительную сумму.")
				continue
			}
			atm.Balance += amount
			fmt.Printf("Успешно пополнено: %.2f руб. Новый баланс: %.2f руб.\n", amount, atm.Balance)

		case "3":
			fmt.Print("Введите сумму для снятия: ")
			amount, err := readAmount(reader)
			if err != nil {
				fmt.Println("Введите корректную положительную сумму.")
				continue
			}
			if amount > atm.Balance {
				fmt.Println("Недостаточно средств на счёте.")
				continue
			}
			atm.Balance -= amount
			fmt.Printf("Успешно снято: %.2f руб. Новый баланс: %.2f руб.\n", amount, atm.Balance)

		case "4":
			fmt.Println("Спасибо за использование банкомата. До свидания!")
			return

		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}
