package main

import (
	"fmt"
	"strings"
)

func main() {
	var num int
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Ошибка при чтении числа:", err)
		return
	}

	if num >= 12307 {
		fmt.Printf("Ваше число: %d. Вычисления не нужны!", num)
		return
	}

	for num < 12307 {
		if num < 0 {
			num *= -1
		} else if num%7 == 0 {
			num *= 39
		} else if num%9 == 0 {
			num = num*13 + 1
			continue
		} else {
			num = (num + 2) * 3
		}

		if num%13 == 0 && num%9 == 0 {
			fmt.Print("service error\n")
			break
		} else {
			num++
		}
	}

	fmt.Printf("Результат вычислений: %d (%s). Ура!", num, numberToString(num))
}

func numberToString(num int) string {
	thousands := []string{"тысяча", "тысячи", "тысяч"}

	if num == 0 {
		return "ноль"
	}

	result := []string{}

	if num >= 1000 {
		numberThousands := num / 1000
		num %= 1000

		var thousandText string
		switch numberThousands {
		case 1:
			thousandText = thousands[0]
		case 2, 3, 4:
			thousandText = thousands[1]
		default:
			thousandText = thousands[2]
		}

		result = append(result, convertThreeDigit(numberThousands), thousandText)
	}

	result = append(result, convertThreeDigit(num))

	return strings.TrimSpace(strings.Join(result, " "))
}

func convertThreeDigit(num int) string {
	if num == 0 {
		return ""
	}

	units := []string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"}
	teens := []string{"десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	tens := []string{"", "", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
	hundreds := []string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}

	result := []string{}

	if num >= 100 {
		result = append(result, hundreds[num/100])
		num %= 100
	}

	if num >= 20 {
		result = append(result, tens[num/10])
		if num%10 != 0 {
			result = append(result, units[num%10])
		}
	} else if num >= 10 {
		result = append(result, teens[num-10])
	} else if num > 0 {
		result = append(result, units[num])
	}

	return strings.Join(result, " ")
}
