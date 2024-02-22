Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
В функции test отложенная функция может считывать
и присваивать именованные возвращаемые значения возвращаемой функции
В функции anotherTest в отложенной функции переменная x увеличивается на 1, 
но это не влияет на x из основной функции, defer не имеет доступа к переменной x
```
