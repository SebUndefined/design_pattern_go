package main

import "fmt"

// Il est possible de retourner seulement un interface
// Cela évite de modifier l'objet et de juste intéragir avec les méthodes de l'interface.

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s and I am %d", p.name, p.age)
}

type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Hi, my name is %s and I am %d and I am old", p.name, p.age)
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name: name, age: age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("Seb", 32)
	p.SayHello()

	pt := NewPerson("Jacky", 102)
	pt.SayHello()
}
