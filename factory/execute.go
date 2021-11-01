package factory

import "fmt"

func printDetails(gun iGun) {
	fmt.Printf("Gun : %s", gun.getName())
	fmt.Println()
	fmt.Printf("Power : %d", gun.getPower())
	fmt.Println()
}

func execute() {
	ak47, _ := getGun("ak47")
	m4a1, _ := getGun("m4a1")

	printDetails(ak47)
	printDetails(m4a1)
}
