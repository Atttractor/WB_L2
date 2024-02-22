package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type State interface {
	work()
}

type machine struct {
	state State
}

func (m *machine) request() {
	m.state.work()
}

func (m *machine) setState(state State) {
	m.state = state
}

type stateWork struct{}

func (s1 *stateWork) work() {
	fmt.Println("machine is working")
}

type stateNotWork struct{}

func (s1 *stateNotWork) work() {
	fmt.Println("machine is not working")
}

//func main() {
//	m := machine{}
//	m.setState(&stateWork{})
//	m.request()
//
//	m.setState(&stateNotWork{})
//	m.request()
//}
