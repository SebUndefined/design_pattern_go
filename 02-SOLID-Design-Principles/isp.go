package main

// ISP
// Simple: Do not put too much in an interface
// Imagine an actual printer that is able to print, scan and fax a document
// Let's do a document

type Document struct {
}

// Let's do a Machine interface that should be able to print, scan and fax a document

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultifunctionPrinter struct {
}

func (m MultifunctionPrinter) Print(d Document) {
	panic("implement me")
}

func (m MultifunctionPrinter) Fax(d Document) {
	panic("implement me")
}

func (m MultifunctionPrinter) Scan(d Document) {
	panic("implement me")
}

// Imagine an old printer that is only able print

type OldFashionedPrinter struct {
}

func (o OldFashionedPrinter) Print(d Document) {
	panic("implement me")
}

func (o OldFashionedPrinter) Fax(d Document) {
	panic("Operation not supported")
}

func (o OldFashionedPrinter) Scan(d Document) {
	panic("Operation not supported")
}

// ISP we separate all the interfaces

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Fax interface {
	Fax(d Document)
}

// A normal printer can only print...

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {

}

// A photocopier can scan and print

type Photocopier struct {
}

func (m Photocopier) Scan(d Document) {

}

func (m Photocopier) Print(d Document) {

}

// In go we can combine interface

type MultifunctionDevice interface {
	Printer
	Scanner
	Fax
}

// Out of subject
// Decorator pattern

type MultifunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultifunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultifunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {

}
