package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
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

// Хотим добавлять новые методы для сотрудников.
type Employee interface {
	getType() string
	accept(Visitor)
}

// Первый тип сотрудника - дизайнер.
type Designer struct {
	id          int
	name        string
	salary      int
	department  string
	currentTask string
}

// Принимаем новый метод.
func (d *Designer) accept(v Visitor) {
	v.visitForDesigner(d)
}

func (d *Designer) getType() string {
	return "Designer"
}

// Второй тип сотрудника - HR.
type HR struct {
	id     int
	name   string
	salary int
}

func (h *HR) accept(v Visitor) {
	v.visitForHR(h)
}

func (h *HR) getType() string {
	return "HR"
}

// Готовы принять методы любой структуры, которая реализует интерфейс Visitor.
type Visitor interface {
	visitForDesigner(*Designer)
	visitForHR(*HR)
}

// Пример структуры, которая реализует Visitor.
type salaryCalculator struct {
	salary int
}

func (s *salaryCalculator) visitForDesigner(d *Designer) {
	if d.name == "Petya" {
		s.salary = 10000

	}
	fmt.Printf("New designer salary: %d\n", s.salary)
}

func (s *salaryCalculator) visitForHR(h *HR) {
	if h.name == "Vasya" {
		s.salary = 10101

	}
	fmt.Printf("New HR salary: %d\n", s.salary)
}

/*
func main() {
	designer1 := &Designer{name: "Petya", salary: 100}
	hr1 := &HR{name: "Vasya", salary: 95}

	salaryCalculator := &salaryCalculator{}

	designer1.accept(salaryCalculator)
	hr1.accept(salaryCalculator)
}
*/
