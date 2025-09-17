package game

import (
	"fmt"
	"strings"
)

type Character struct {
	Nom                string
	Classe             string
	Niveau             int
	PVMax              int
	PVActuels          int
	Inventaire         []string
	Skill              []string
	Argent             int
	Equipment          Equipment
	CapaciteInventaire int
}

func Creation() Character {
	var nom, classe string

	fmt.Println("Choisissez le nom de votre personnage :")
	fmt.Scan(&nom)
	if len(nom) > 0 {
		nom = strings.ToUpper(nom[:1]) + strings.ToLower(nom[1:])
	}

	fmt.Println("Choisissez votre classe (Humain / Elfe / Nain) :")
	fmt.Scan(&classe)

	pvMax := 100
	switch strings.ToLower(classe) {
	case "elfe":
		pvMax = 80
	case "nain":
		pvMax = 120
	case "humain":
		pvMax = 100
	default:
		fmt.Println("Classe invalide, vous serez Humain par défaut.")
		classe = "Humain"
		pvMax = 100
	}

	return Character{
		Nom:                nom,
		Classe:             classe,
		Niveau:             1,
		PVMax:              pvMax,
		PVActuels:          pvMax,
		Inventaire:         []string{},
		Skill:              []string{"Coup de poing"},
		Argent:             100,
		Equipment:          Equipment{},
		CapaciteInventaire: 10,
	}
}

func DisplayInfo(c Character) {
	fmt.Println("Nom :", c.Nom)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Niveau)
	fmt.Println("Points de vie :", c.PVActuels, "/", c.PVMax)
	fmt.Println("Inventaire :", c.Inventaire)
}

func IsDead(c *Character) bool {
	if c.PVActuels <= 0 {
		fmt.Println(c.Nom, "est mort!")
		c.PVActuels = c.PVMax / 2
		fmt.Println(c.Nom, "a été ressuscité avec", c.PVActuels, "PV.")
		return true
	}
	return false
}

func SpellBook(c *Character) {
	for _, s := range c.Skill {
		if s == "Boule de feu" {
			fmt.Println("Vous connaissez déjà ce sort.")
			return
		}
	}
	c.Skill = append(c.Skill, "Boule de feu")
	fmt.Println("Vous avez appris le sort : Boule de feu")
}
