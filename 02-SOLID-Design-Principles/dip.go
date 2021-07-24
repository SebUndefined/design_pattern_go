package main

import "fmt"

// Dependency Inversion principle
// HLM should not depend on LLM
// Both should depend on abstractions

// Let's imagine we want to model relationship between people (parent, child etc...)

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// Other attribute
}

type Info struct {
	from         *Person
	relationShip Relationship
	to           *Person
}

// Low level module

type Relationships struct {
	relations []Info
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationShip == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Relationships) AddParentToChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// High level module

type Research struct {
	// Break DIP
	//relationships Relationships
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	/* Old code
	relations := r.relationships.relations

	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationShip == Parent {
			fmt.Println("John has a child called ", rel.to.name)
		}
	}*/

	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child named ", p.name)
	}
}

func main() {
	parent := Person{"John"}

	child1 := Person{"Chris"}
	child2 := Person{"Sébastien"}

	relation := Relationships{}
	relation.AddParentToChild(&parent, &child1)
	relation.AddParentToChild(&parent, &child2)

	r := Research{&relation}
	r.Investigate()
}
