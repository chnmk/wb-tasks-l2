Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

Функция test возвращает указатель, который реализует встроенный интерфейс error, даже если является nil.

Сравнение err != nil происходит и по значению, и по типу. В данном случае тип не является nil (у интерфейса под капотом присутствует поле itab).
```
