package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
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

// Интерфейс, который должен быть реализован всеми обработчиками в цепочкею
type Handler interface {
	handle(*Request)
	setNext(Handler)
}

// Первый элемент цепочки.
type Validator struct {
	next Handler
}

func (h *Validator) handle(r *Request) {
	if r.Type == 3 {
		fmt.Println("Can't validate request of type 3")
	} else {
		fmt.Println("Validating...")
	}

	h.next.handle(r)
}

func (r *Validator) setNext(next Handler) {
	r.next = next
}

// Второй элемент цепочки.
type Converter struct {
	next Handler
}

func (h *Converter) handle(r *Request) {
	if r.Type == 3 {
		fmt.Println("Can't convert request of type 3")
	} else {
		fmt.Println("Converting...")
	}

	h.next.handle(r)
}

func (r *Converter) setNext(next Handler) {
	r.next = next
}

// Третий элемент цепочки.
type Printer struct {
	next Handler
}

func (h *Printer) handle(r *Request) {
	if r.Type == 2 {
		fmt.Println("Can't print request of type 2")
	} else {
		fmt.Println(r.Content)
	}
}

func (r *Printer) setNext(next Handler) {
	r.next = next
}

// Запрос, который нужно обработать.
type Request struct {
	Type    int
	Content string
}

/*
func main() {
	validator := &Validator{}

	converter := &Converter{}
	validator.setNext(converter)

	printer := &Printer{}
	converter.setNext(printer)

	request := &Request{Type: 3, Content: "Hello"}

	validator.handle(request)
}
*/
