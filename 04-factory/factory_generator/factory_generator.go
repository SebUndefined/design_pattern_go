package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// Functionnal

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// Other

type EmployeeFactory struct {
	Position    string
	AnnalIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnalIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {

	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 100000)

	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	fmt.Println(developer)
	fmt.Println(manager)

	bossFactory := NewEmployeeFactory2("boss", 200000)
	boss := bossFactory.Create("John")
	fmt.Println(boss)
}
