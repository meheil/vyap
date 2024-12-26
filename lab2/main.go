package main

import (
	"errors"
	"fmt"
	"math"
)

// Задание 1. Массивы и срезы

// formatIP принимает IP-адрес в виде массива из четырех байтов и возвращает строку в формате "127.0.0.1".
func formatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// listEven принимает на вход диапазон и возвращает срез с четными числами и значение типа error.
func listEven(start, end int) ([]int, error) {
	if start > end {
		return nil, errors.New("левая граница диапазона больше правой")
	}

	var evens []int
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	return evens, nil
}

// Задание 2. Карты

// countCharacters принимает строку и возвращает карту с количеством вхождений каждого символа.
func countCharacters(s string) map[string]int {
	charCount := make(map[string]int)
	for _, char := range s {
		charCount[string(char)]++
	}
	return charCount
}

// Задание 3. Структуры, методы и интерфейсы

// Определяем структуру Point (точка)
type Point struct {
	X, Y float64
}

// Определяем структуру Line (отрезок)
type Line struct {
	Start, End Point
}

// Метод для структуры Line, возвращающий длину отрезка
func (l Line) Length() float64 {
	return math.Sqrt(math.Pow(l.End.X-l.Start.X, 2) + math.Pow(l.End.Y-l.Start.Y, 2))
}

// Определяем структуру Triangle (треугольник)
type Triangle struct {
	A, B, C Point
}

// Метод для структуры Triangle, возвращающий площадь треугольника
func (t Triangle) Area() float64 {
	// Используем формулу Герона для вычисления площади треугольника
	a := Line{t.A, t.B}.Length()
	b := Line{t.B, t.C}.Length()
	c := Line{t.C, t.A}.Length()
	s := (a + b + c) / 2 // Полупериметр
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

// Определяем структуру Circle (круг)
type Circle struct {
	Center Point
	Radius float64
}

// Метод для структуры Circle, возвращающий площадь круга
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Определяем интерфейс Shape с единственным методом Area
type Shape interface {
	Area() float64
}

// Функция для вывода площади фигуры, принимающая интерфейс Shape
func printArea(s Shape) {
	result := s.Area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}

// Функция высшего порядка Map, которая применяет переданную функцию к каждому элементу среза
func Map(slice []float64, fn func(float64) float64) []float64 {
	// Создаем новый срез для хранения результатов
	result := make([]float64, len(slice))
	copy(result, slice) // Копируем исходный срез, чтобы сохранить его неизменным

	for i, v := range result {
		result[i] = fn(v) // Применяем функцию ко всем элементам среза
	}
	return result
}

// Функция для возведения числа в квадрат
func square(x float64) float64 {
	return x * x
}

func main() {
	// Пример использования formatIP
	ip := [4]byte{192, 168, 1, 1}
	fmt.Println("Formatted IP:", formatIP(ip))

	// Пример использования listEven
	evens, err := listEven(1, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Even numbers:", evens)
	}

	// Пример использования countCharacters
	text := "hello world"
	fmt.Println("Character counts:", countCharacters(text))

	// Создаем экземпляр треугольника и выводим его площадь
	triangle := Triangle{
		A: Point{0, 0},
		B: Point{4, 0},
		C: Point{4, 3},
	}
	printArea(triangle)

	// Создаем экземпляр круга и выводим его площадь
	circle := Circle{
		Center: Point{0, 0},
		Radius: 5,
	}
	printArea(circle)

	// Создаем срез с исходными значениями
	values := []float64{1, 2, 3, 4, 5}

	// Применяем функцию Map для возведения в квадрат каждого элемента среза
	// Передаем срез и функцию square в качестве аргументов
	squaredValues := Map(values, square)

	// Выводим исходный срез и результат
	fmt.Println("Original slice:", values)
	fmt.Println("Squared slice:", squaredValues)
}
