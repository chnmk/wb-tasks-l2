package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
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

// Клиент, которому мы хотим посылать сигналы.
type Client interface {
	start()
	shutdown()
}

type MyClient struct {
	active bool
}

func (c *MyClient) start() {
	c.active = true
	fmt.Println("Starting...")
}

func (c *MyClient) shutdown() {
	c.active = false
	fmt.Println("Shutting down...")
}

// Хотим создать объекты, которые отправляли бы определенный сигнал клиенту.
type Switch struct {
	command Command
}

func (p *Switch) switchSignal() {
	p.command.do()
}

// Создаём разные виды команд, чтобы использовать их в разных объектах.
type Command interface {
	do()
}

// Команда для запуска клиента.
type StartCommand struct {
	client Client
}

func (c *StartCommand) do() {
	c.client.start()
}

// Команда для отключения клиента.
type ShutdownCommand struct {
	client Client
}

func (c *ShutdownCommand) do() {
	c.client.shutdown()
}

/*
func main() {
	myClient := &MyClient{}

	startCommand := &StartCommand{client: myClient}
	shutdownCommand := &ShutdownCommand{client: myClient}

	onSwitch := &Switch{command: startCommand}
	onSwitch.switchSignal()

	go func() {
		onSwitch.switchSignal()
	}()

	time.Sleep(1 * time.Second)

	offSwitch := &Switch{command: shutdownCommand}
	offSwitch.switchSignal()
}
*/
