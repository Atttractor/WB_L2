package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Context struct {
	Strategy
}

func (c *Context) setStrategy(strategy Strategy) {
	c.Strategy = strategy
}

type Strategy interface {
	Func()
}

type Strategy1 struct{}

func (s *Strategy1) Func() {
	fmt.Println("Strategy1")
}

type Strategy2 struct{}

func (s *Strategy2) Func() {
	fmt.Println("Strategy2")
}

//func main() {
//	ctx := Context{}
//	ctx.setStrategy(new(Strategy1))
//	ctx.Strategy.Func()
//	ctx.setStrategy(new(Strategy2))
//	ctx.Strategy.Func()
//}
