package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

// factory function est juste une fonction qui permet de retourner un objet
// en fonction des arguments de la méthode

func NewPerson(name string, age int) *Person {
	if age < 16 {
		//...
	}
	return &Person{name, age, 2}
}

func main() {
	p := NewPerson("Seb", 32)
	fmt.Println(p)

}
