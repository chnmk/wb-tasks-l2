package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Строитель используется для пошагового создания объектов.

	Плюсы:
		- Можно переиспользовать для создания разных объектов
		- Упрощает работу с сложными объектами, скрывает сложную логику где это необходимо
	Минусы:
		- При использовании с простыми объектами может оказаться лишней надстройкой
		- Необходимо следить за последовательностью создания объектов при использовании разных конкретных строителей на один абстрактный

	Примеры использования:
		- Последовательное создание сложных объектов в соответствии с бизнес-логикой
*/

// Хотим записать день рождения пользователя в определенном формате.
type User struct {
	Birthday string
}

type AbstractBuilder interface {
	SetDay(string)
	SetMonth(string)
	SetYear(string)
	GetUser() User
}

// Конкретный строитель для получения информации о пользователе:
type UserBuilder struct {
	User User
}

// Директор отдаёт команды строителю, определяет порядок исполнения.
type Director struct {
	builder AbstractBuilder
}

func (d *Director) Construct() User {
	d.builder.SetDay("01")
	d.builder.SetMonth("January")
	d.builder.SetYear("2000")

	return d.builder.GetUser()
}

func (u *UserBuilder) SetDay(input string) {
	u.User.Birthday = input
}

func (u *UserBuilder) SetMonth(input string) {
	u.User.Birthday = u.User.Birthday + " " + input
}

func (u *UserBuilder) SetYear(input string) {
	u.User.Birthday = u.User.Birthday + ", " + input
}

func (u *UserBuilder) GetUser() User {
	return u.User
}

/*
func main() {
	var D Director
	D.builder = &UserBuilder{}

	user := D.Construct()

	fmt.Println(user.Birthday)
}
*/
