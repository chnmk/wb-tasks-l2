package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Switch struct {
	command Command
}

func (p *Switch) switchSignal() {
	p.command.execute()
}

type Command interface {
	execute()
}

type StartCommand struct {
	client Client
}

func (c *StartCommand) execute() {
	c.client.start()
}

type ShutdownCommand struct {
	client Client
}

func (c *ShutdownCommand) execute() {
	c.client.shutdown()
}

type Client interface {
	start()
	shutdown()
}

type MyClient struct {
	isRunning bool
}

func (c *MyClient) start() {
	c.isRunning = true
	fmt.Println("Starting...")
}

func (c *MyClient) shutdown() {
	c.isRunning = false
	fmt.Println("Shutting down...")
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
