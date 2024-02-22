package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Service interface {
	execute()
	setNext(Service)
}

type chain1 struct {
	next Service
}

func (c1 *chain1) execute() {
	fmt.Println("executing chain 1")
	c1.next.execute()
}

func (c1 *chain1) setNext(nextChain Service) {
	c1.next = nextChain
}

type chain2 struct {
	next Service
}

func (c2 *chain2) execute() {
	fmt.Println("executing chain 2")
	c2.next.execute()
}

func (c2 *chain2) setNext(nextChain Service) {
	c2.next = nextChain
}

type chain3 struct {
	next Service
}

func (c3 *chain3) execute() {
	fmt.Println("executing chain 3")
}

func (c3 *chain3) setNext(nextChain Service) {
	c3.next = nextChain
}

//func main() {
//	c1 := chain1{}
//	c2 := chain2{}
//	c3 := chain3{}
//	c1.setNext(&c2)
//	c2.setNext(&c3)
//	c1.execute()
//}
