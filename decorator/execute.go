package decorator

import "fmt"

func execute() {
	pizza := &veggieMania{}

	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with cheese and tomato topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
