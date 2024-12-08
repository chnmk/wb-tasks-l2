package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Фасад используется для предоставления общего интерфейса к нескольким связанным подсистемам.

	Плюсы:
		- Упрощает работу с сложными системами
		- Скрывает сложную логику где это необходимо
		- Понятен, легко реализуется и используется
	Минусы:
		- При использовании с простыми системами может оказаться лишней надстройкой
		- При использовании с слишком сложными системами существует вероятность перегрузки фасада

	Примеры использования:
		- Обеспечение доступа к методам фреймворка или другого отдельного от основной логики сервиса инструмента
		- Обеспечение доступа к схожим методам разных структур
*/

// Некие источники данных, к которым хотелось бы иметь общий доступ
var Source1 DataSource1
var Source2 DataSource2
var Source3 DataSource3

type DataSource1 struct {
	data string
}

type DataSource2 struct {
	data string
}

type DataSource3 struct {
	data string
}

func (d *DataSource1) Fetch() {
	// Какая-то логика получения данных
	// ...

	d.data = "aaa"
}

func (d *DataSource2) Fetch() {
	// Какая-то логика получения данных
	// ...

	d.data = "bbb"
}

func (d *DataSource3) Fetch() {
	// Какая-то логика получения данных
	// ...

	d.data = ""
}

// Фасад для получения данных из всех источников
var fetchFacade FetchFacade

type FetchFacade struct {
	FullData []string
}

func (f *FetchFacade) FetchAllSources() []string {
	var fullData []string

	Source1.Fetch()
	Source2.Fetch()
	Source3.Fetch()

	fullData = append(fullData, Source1.data, Source2.data, Source3.data)

	f.FullData = fullData

	return fullData
}

/*
func main() {
	d := fetchFacade.FetchAllSources()
	fmt.Println(d)
}
*/
