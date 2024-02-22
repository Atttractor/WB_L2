package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type command interface {
	execute()
}

type receiver struct{}

func (r *receiver) action(text string) {
	fmt.Println(text)
}

type command1 struct {
	receiver *receiver
}

func (c1 *command1) execute() {
	c1.receiver.action("command 1")
}

type command2 struct {
	receiver *receiver
}

func (c2 *command2) execute() {
	c2.receiver.action("command 2")
}

type invoker struct {
	masOfCommands []command
}

func (i *invoker) add(command command) {
	i.masOfCommands = append(i.masOfCommands, command)
}

func (i *invoker) executeCommands() {
	for _, command := range i.masOfCommands {
		command.execute()
	}
}

//func main() {
//	r := receiver{}
//	inv := invoker{
//		masOfCommands: []command{
//			&command1{receiver: &r},
//			&command2{receiver: &r},
//			&command1{receiver: &r}},
//	}
//	inv.executeCommands()
//}
