package main

import "fmt"

type equipment interface {
	Accept(visitor)
}

type hdd struct {
	price int
}

func (e *hdd) accept(visitor visitor) {
	visitor.visitHDD(e)
}

type ram struct {
	price int
}

func (e *ram) accept(visitor visitor) {
	visitor.visitRAM(e)
}

type visitor interface {
	visitHDD(*hdd)
	visitRAM(*ram)
}

type priceVisitor struct {
	total int
}

func (v *priceVisitor) visitHDD(element *hdd) {
	// fmt.Println("ConcreteVisitor.VisitA()")
	v.total += element.price
	fmt.Printf("ConcreteVisitor.VisitHDD() total %d \n", v.total)
}

func (v *priceVisitor) visitRAM(element *ram) {
	v.total += element.price
	fmt.Printf("ConcreteVisitor.VisitRAM() total %d \n", v.total)

}

func main() {
	visitor := new(priceVisitor)
	hdd := new(hdd)
	ram := new(ram)

	hdd.price = 10
	ram.price = 10
	hdd.accept(visitor)
	ram.accept(visitor)
}
