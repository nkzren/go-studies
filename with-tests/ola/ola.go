package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

func Ola(name string) string {
	return prefixoOlaPortugues + name
}

func main() {
	fmt.Println(Ola("mundo"))
}
