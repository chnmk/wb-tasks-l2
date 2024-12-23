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
2, 1

defer выполняется после вычисления возвращаемых переменных. В test() возвращается переменная x, который после этого успевает измениться через defer. В anotherTest() возвращается псевдопеременная, которая копирует значение x до defer, и дальнейшие изменения x на неё не влияют.

https://go.dev/ref/spec#Defer_statements

For instance, if the deferred function is a function literal and the surrounding function has named result parameters that are in scope within the literal, the deferred function may access and modify the result parameters before they are returned. If the deferred function has any return values, they are discarded when the function completes.
```
