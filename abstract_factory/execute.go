package abstractfactory

import "fmt"

func Execute() {
	adidasFactory, _ := getSportFactory("adidas")
	nikeFactory, _ := getSportFactory("nike")

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	fmt.Println("Adidas Shoe: ", adidasShoe, "Adidas Shirt: ", adidasShirt)
	fmt.Println("Nike Shoe: ", nikeShoe, "Nike Shirt: ", nikeShirt)
}
