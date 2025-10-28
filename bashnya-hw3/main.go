package main

import "fmt"

func main() {
	stack := NewStack()

	stack.Push(20)
	stack.Push(35)

	if value, err := stack.Pop(); err == nil {
		fmt.Printf("Извлечено: %d\n", value)
	} else {
		fmt.Printf("Ошибка: %v\n", err)
	}
}
