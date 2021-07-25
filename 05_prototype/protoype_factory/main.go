package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite                        int
	StreetAddress, City, Country string
}

type Employee struct {
	Name   string
	Office Address
}

func (p *Employee) DeepCopy() *Employee {
	// Créer un buffer
	b := bytes.Buffer{}
	// Créer un encoder avec le pointeur du buffer
	e := gob.NewEncoder(&b)
	// Encoder Person
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	// Créer un décodeur
	d := gob.NewDecoder(&b)
	// Préparer la mémoire
	result := Employee{}
	// Décoder en passant le pointeur de la variable result
	_ = d.Decode(&result)
	return &result
}

var mainOffice = Employee{
	"", Address{0, "123 East Dr", "London", "UK"},
}
var auxOffice = Employee{
	"", Address{0, "66 West Dr", "London", "UK"},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}
func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
