package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

func newBuilder() builderObject {
	return builderObject{obj: object{}}
}

type object struct {
	field1 string
	field2 int
	field3 bool
}

type builderObject struct {
	obj object
}

func (bo *builderObject) setField1(data string) {
	bo.obj.field1 = data
}

func (bo *builderObject) setField2(data int) {
	bo.obj.field2 = data
}

func (bo *builderObject) setField3(data bool) {
	bo.obj.field3 = data
}

func (bo *builderObject) build() object {
	return bo.obj
}

func main() {
	builder := newBuilder()
	builder.setField1("123")
	builder.setField2(123)
	builder.setField3(false)
	fmt.Println(builder.build())
}
