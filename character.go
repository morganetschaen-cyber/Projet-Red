package main

import "fmt"
import "strings"

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PVActuels  int
	Inventaire []string
	Skill      []string
	Argent     int
	Equipment	Equipment
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
		Argent:     100
		Equipment: Equipment{},
	}
}

func characterCreation() Character {
	var nom string
	var classe string

	fmt.Println("Choisissez le nom de votre personnage :")
	fmt.Scan(&nom)

	if len(nom) > 0 {
		nom = strings.ToUpper(nom[:1]) + strings.ToLower(nom[1:])
	}

	fmt.Println("Choisissez votre classe (Humain / Elfe / Nain) :")
	fmt.Scan(&classe)

	var pvMax int
	switch strings.ToLower(classe) {
	case "humain":
		pvMax = 100
	case "elfe":
		pvMax = 80
	case "nain":
		pvMax = 120
	default:
		fmt.Println("Classe invalide, par défaut vous serez Humain.")
		pvMax = 100
		classe = "Humain"
	}

	return Character{
		Nom:        nom,
		Classe:     classe,
		PVMax:      pvMax,
		PVActuels:  pvMax / 2,
		Niveau:     1,
		Inventaire: []string{},
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
