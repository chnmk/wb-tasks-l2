package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
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

type State interface {
	describe()
}

// Структура человека с различными состояниями.
type Human struct {
	name          string
	currentState  State
	workingState  State
	studyingState State
	idlingState   State
}

func (p *Human) describe() {
	p.currentState.describe()
}

// Установка состояния для человека.
func (p *Human) setState(state State) {
	p.currentState = state
}

// Первое состояние и функция, которая будет вызываться при этом состоянии.
type WorkingState struct {
	human *Human
}

func (s *WorkingState) describe() {
	fmt.Printf("%s is working\n", s.human.name)
}

// Второе состояние, которая будет вызываться при этом состоянии.
type StudyingState struct {
	human *Human
}

func (s *StudyingState) describe() {
	fmt.Printf("%s is studying\n", s.human.name)
}

// Третье состояние, которая будет вызываться при этом состоянии.
type IdlingState struct {
	human *Human
}

func (s *IdlingState) describe() {
	fmt.Printf("%s is't doing anything\n", s.human.name)
}

/*
func main() {
	human1 := &Human{name: "Petya"}

	workingState := &WorkingState{human: human1}
	studyingState := &StudyingState{human: human1}
	idlingState := &IdlingState{human: human1}

	human1.workingState = workingState
	human1.studyingState = studyingState
	human1.idlingState = idlingState

	human1.setState(idlingState)
	human1.describe()
	human1.setState(workingState)
	human1.describe()
}
*/
