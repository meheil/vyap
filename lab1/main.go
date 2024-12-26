package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Напишите Ваше имя")
	var name string
	fmt.Scan(&name)
	hello(name)

	fmt.Println("Тестирование функции printEven:")
	if err := printEven(3, 10); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Чётные числа от 3 до 10:")
	}

	if err := printEven(10, 20); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Чётные числа от 10 до 20:")
	}

	if err := printEven(5, 3); err != nil {
		fmt.Println(err)
	}

	fmt.Println()

	fmt.Println("Тестирование функции apply:")
	result, err := apply(3, 5, "+")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("3 + 5 =", result)
	}

	result, err = apply(7, 10, "*")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("7 * 10 =", result)
	}

	result, err = apply(3, 5, "#")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("3 # 5 =", result)
	}

	result, err = apply(10, 0, "/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("10 / 0 =", result)
	}
}

func hello(name string) {
	fmt.Println("Привет, " + name + "!")
}

func printEven(a, b int64) error {
	if a > b {
		return errors.New("левая граница больше правой")
	}
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	return nil
}

func apply(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("действие не поддерживается")
	}
}
