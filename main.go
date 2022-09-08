package main

import abstractfactory "design_patterns/abstract_factory"

func main() {
	adidasFactory, _ := abstractfactory.GetSportsFactory("adidas")
	nikeFactory, _ := abstractfactory.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	abstractfactory.PrintShoeDetails(nikeShoe)
	abstractfactory.PrintShirtDetails(nikeShirt)

	abstractfactory.PrintShoeDetails(adidasShoe)
	abstractfactory.PrintShirtDetails(adidasShirt)
}
