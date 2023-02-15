package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

func Ola(name string) string {
	if name == "" {
		name = "mundo"
	}
	return prefixoOlaPortugues + name
}

func main() {
	fmt.Println(Ola("mundo"))
}
