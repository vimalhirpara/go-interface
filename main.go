package main

import "fmt"

type Printer interface {
	Print() string
}

type myPrinter struct {
}

func (mp myPrinter) Print() string {
	return "Printing one page..."
}

type secondPrinter struct {
}

func (mp secondPrinter) Print() string {
	return "Printing five pages..."
}

func process(quipment Printer) {
	fmt.Println("Running print...", quipment.Print())
}

func main() {
	printer := myPrinter{}
	process(printer)

	secondprinter := secondPrinter{}
	process(secondprinter)
}
