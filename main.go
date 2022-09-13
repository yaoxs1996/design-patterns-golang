package main

import (
	abstractfactory "design_patterns/abstract_factory"
	builder "design_patterns/builder"
	factorymethod "design_patterns/factory_method"
	"fmt"
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

	// 生成器模式
	fmt.Println("<===== 生成器模式 =====>")
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()
	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)

}
