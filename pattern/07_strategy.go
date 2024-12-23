package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Описание...

	Плюсы:
		- ...
		- ...
	Минусы:
		- ...

	Примеры использования:
		- ...
*/

// Интерфейс, объединяющий различные варианты похожих функций.
type NumPrinter interface {
	print(n *Numbers)
}

// Структура для первого варианта функции.
type PrintTo10 struct {
}

func (l *PrintTo10) print(n *Numbers) {
	for i, n := range n.Content {
		if n <= 10 {
			fmt.Printf("%d: %d\n", i, n)
		}
	}
}

// Структура для второго варианта функции.
type PrintTo20 struct {
}

func (l *PrintTo20) print(n *Numbers) {
	for i, n := range n.Content {
		if n <= 20 {
			fmt.Printf("%d: %d\n", i, n)
		}
	}
}

// Структура, на которой хотим использовать методы.
type Numbers struct {
	Content    []int
	NumPrinter NumPrinter
}

func getNumbers() *Numbers {
	return &Numbers{Content: []int{1, 3, 6, 10, 15, 19}}
}

// Выбор метода для структуры.
func (n *Numbers) setNumPrinter(p NumPrinter) {
	n.NumPrinter = p
}

func (n *Numbers) print() {
	n.NumPrinter.print(n)
}

/*
func main() {
	numbers := getNumbers()

	printTo10 := &PrintTo10{}
	numbers.setNumPrinter(printTo10)
	numbers.print()

	fmt.Println("")

	printTo20 := &PrintTo20{}
	numbers.setNumPrinter(printTo20)
	numbers.print()
}
*/
