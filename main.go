package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	operations := map[string]func(float64, float64) float64{
		"1": func(a, _ float64) float64 { return math.Sin(a * math.Pi / 180) },
		"2": func(a, _ float64) float64 { return math.Cos(a * math.Pi / 180) },
		"3": func(a, _ float64) float64 { return math.Tan(a * math.Pi / 180) },
		"4": func(a, _ float64) float64 { return 1 / math.Tan(a*math.Pi/180) },
		"5": func(a, b float64) float64 { return a + b },
		"6": func(a, b float64) float64 { return a - b },
		"7": func(a, b float64) float64 { return a * b },
		"8": func(a, b float64) float64 { return a / b },
	}

	for {
		fmt.Println("Калькулятор")
		fmt.Println("1. Синус")
		fmt.Println("2. Косинус")
		fmt.Println("3. Тангенс")
		fmt.Println("4. Котангенс")
		fmt.Println("5. Сложение")
		fmt.Println("6. Вычитание")
		fmt.Println("7. Умножение")
		fmt.Println("8. Деление")
		fmt.Println("9. Выход")

		fmt.Print("Введите первое число: ")
		var num1 string
		for {
			fmt.Scanln(&num1)

			if n1, err := strconv.ParseFloat(num1, 64); err != nil {
				fmt.Println("Не вводите буквы, введите число")
			} else {
				fmt.Println("Выберите операцию:")
				var choice string
				for {
					fmt.Scanln(&choice)

					if op, ok := operations[choice]; ok {
						if choice == "5" || choice == "6" || choice == "7" || choice == "8" {
							fmt.Print("Введите второе число: ")
							var num2 string
							for {
								fmt.Scanln(&num2)

								if n2, err := strconv.ParseFloat(num2, 64); err != nil {
									fmt.Println("Не вводите буквы, введите число")
								} else {
									fmt.Printf("%.2f %s %.2f = %.2f\n", n1, getOperationName(choice), n2, op(n1, n2))
									break
								}
							}
						} else {
							fmt.Printf("%.2f %s = %.2f\n", n1, getOperationName(choice), op(n1, 0))
						}
						break
					} else if choice == "9" {
						return
					} else {
						fmt.Println("Неправильный выбор. Введите номер операции.")
					}
				}
				break
			}
		}

		fmt.Print("Хотите продолжить использовать калькулятор? (да/нет): ")
		var answer string
		for {
			fmt.Scanln(&answer)

			if answer == "да" || answer == "y" {
				break
			} else if answer == "нет" || answer == "n" {
				return
			} else {
				fmt.Println("Неправильный ввод. Пожалуйста, введите 'да' или 'нет'.")
			}
		}
	}
}

func getOperationName(choice string) string {
	switch choice {
	case "1":
		return "Син"
	case "2":
		return "Кос"
	case "3":
		return "Тан"
	case "4":
		return "Кот"
	case "5":
		return "+"
	case "6":
		return "-"
	case "7":
		return "*"
	case "8":
		return "/"
	default:
		return ""
	}
}
