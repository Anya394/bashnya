package main

import "fmt"

func main() {
	var name string
	fmt.Print("Введите ваше имя: ")
	fmt.Scan(&name)
	fmt.Printf("Привет, %s!\n", name)
}
