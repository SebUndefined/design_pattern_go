package main

import "fmt"

type People struct {
	name, position, company string
}

type peopleMod func(*People)
type PeopleBuilder struct {
	actions []peopleMod
}

func (b *PeopleBuilder) Called(name string) *PeopleBuilder {
	b.actions = append(b.actions, func(people *People) {
		people.name = name
	})
	return b
}

func (b *PeopleBuilder) WorkAsA(position string) *PeopleBuilder {
	b.actions = append(b.actions, func(people *People) {
		people.position = position
	})
	return b
}

func (b *PeopleBuilder) At(company string) *PeopleBuilder {
	b.actions = append(b.actions, func(people *People) {
		people.company = company
	})
	return b
}

func (b *PeopleBuilder) Build() *People {
	p := People{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func main() {
	b := PeopleBuilder{}
	p := b.Called("Dimitry").At("Inforsud").WorkAsA("Dev").Build()
	fmt.Println(p)
}
