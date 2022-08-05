package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	j := &Person{"Joao Silva", 42}
	k := &Person{"Kleber Pereira", 9001}
	fmt.Println(j, k)
}
