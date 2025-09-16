package main

import "fmt"

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PVActuels  int
	Inventaire []string
	Skill      []string
}

func initCharacter(nom, classe string, niveau, pvMax, pvActuels int, inventaire []string) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PVActuels:  pvActuels,
		Inventaire: inventaire,
		Skill:      []string{"Coup de poing"},
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

func isDead(c *Character) bool {
	if c.PVActuels <= 0 {
		fmt.Println(c.Nom, "est mort!")
		c.PVActuels = c.PVMax / 2
		fmt.Println(c.Nom, "a été ressuscité avec", c.PVActuels, "PV.")
		return true
	}
	return false
}
func spellBook(c *Character) {
	for _, s := range c.Skill {
		if s == "Boule de feu" {
			fmt.Println("Vous connaissez déjà ce sort.")
			return
		}
	}
	c.Skill = append(c.Skill, "Boule de feu")
	fmt.Println("Vous avez appris le sort : Boule de feu")
}
