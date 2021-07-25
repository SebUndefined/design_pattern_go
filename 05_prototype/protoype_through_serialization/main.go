package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
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
	result := Person{}
	// Décoder en passant le pointeur de la variable result
	_ = d.Decode(&result)
	return &result
}

func main() {
	john := Person{"John", &Address{"123 London RD", "London", "UK"}, []string{"Chris", "Matt"}}
	jane := john.DeepCopy()

	jane.Friends = append(jane.Friends, "Angela")
	jane.Name = "Jane"

	jane.Address.StreetAddress = "321 Baker street"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
