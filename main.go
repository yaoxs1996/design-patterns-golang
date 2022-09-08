package main

import (
	abstractfactory "design_patterns/abstract_factory"
	factorymethod "design_patterns/factory_method"
)

func main() {
	// 工厂方法
	ak47, _ := factorymethod.GetGun("ak47")
	musket, _ := factorymethod.GetGun("musket")
	factorymethod.PrintDetails(ak47)
	factorymethod.PrintDetails(musket)

	// 抽象工厂
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
