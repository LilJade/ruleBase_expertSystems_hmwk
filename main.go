package main

import "fmt"

func main() {
	/*Creating a type People*/
	type People struct {
		name, gender, liking string
	}

	/*Creating a list of people*/
	var list [6]People
	list[0] = People{"Juan", "Hombre", "Comer"}
	list[1] = People{"Pedro", "Hombre", "Bailar"}
	list[2] = People{"Luis", "Hombre", "Bailar"}
	list[3] = People{"Mateo", "Hombre", ""}
	list[4] = People{"Ana", "Mujer", "Comer"}
	list[5] = People{"Maria", "Mujer", ""}

	fmt.Println("A María le agradan los chicos que gustan de bailar.\nEstos chicos son:")

	for _, listing  := range list {
		if listing.gender == "Hombre" && listing.liking == "Bailar" {
			fmt.Println(listing.name)
		}
	}
}
