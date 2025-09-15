package main

import "fmt"

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PVActuels  int
	Inventaire []string
}

func initCharacter(nom, classe string, niveau, pvMax, pvActuels int, inventaire []string) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PVActuels:  pvActuels,
		Inventaire: inventaire,
	}
}

func displayInfo(c Character) {
	fmt.Println("Nom :", c.Nom)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Niveau)
	fmt.Println("Points de vie :", c.PVActuels, "/", c.PVMax)
	fmt.Println("Inventaire :", c.Inventaire)
}

func main() {
	inventaire := []string{"Potion", "Potion", "Potion"}
	heros := initCharacter("Eva", "Elfe", 1, 100, 50, inventaire)

	displayInfo(heros)
}

func takePot(c *Character) {
}
