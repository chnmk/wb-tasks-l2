package pattern

import (
	"errors"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type IProduct interface {
	setName(name string)
	setPrice(price int)
	getName() string
	getPrice() int
}

type Product struct {
	name  string
	price int
}

func (t *Product) setName(name string) {
	t.name = name
}

func (t *Product) getName() string {
	return t.name
}

func (t *Product) setPrice(price int) {
	t.price = price
}

func (t *Product) getPrice() int {
	return t.price
}

type Tablet struct {
	Product
}

func newTablet1() IProduct {
	return &Tablet{
		Product: Product{
			name:  "Tablet-1",
			price: 10000,
		},
	}
}

type Laptop struct {
	Product
}

func newLaptop1() IProduct {
	return &Laptop{
		Product: Product{
			name:  "Laptop-1",
			price: 20000,
		},
	}
}

func getProduct(t string) (IProduct, error) {
	switch t {
	case "Tablet-1":
		return newTablet1(), nil
	case "Laptop-1":
		return newLaptop1(), nil
	default:
		return nil, errors.New("invalid product")
	}

}

/*
func main() {
	t, err := getProduct("Tablet-1")
	if err != nil {
		log.Println(err)
	}

	l, err := getProduct("Laptop-1")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(t.getName())
	fmt.Println(l.getName())
}
*/
