package main

import "fmt"

// Nous donne un objet préconfiguré

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 60000}
	case Manager:
		return &Employee{"", "manager", 100000}
	default:
		panic("Format not supporter")
	}
}

func main() {
	e := NewEmployee(Developer)
	e.Name = "Seb"
	fmt.Println(e)
}
