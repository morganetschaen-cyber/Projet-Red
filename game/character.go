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
	var nom string
	var choixClasse int
	var classe string
	var pvMax int

	fmt.Println("Choisissez le nom de votre personnage :")
	fmt.Scan(&nom)
	if len(nom) > 0 {
		nom = strings.ToUpper(nom[:1]) + strings.ToLower(nom[1:])
	}

	fmt.Println("Choisissez votre classe :")
	fmt.Println("1. Humain (PV 100)")
	fmt.Println("2. Elfe (PV 80)")
	fmt.Println("3. Nain (PV 120)")
	fmt.Print("> ")
	fmt.Scan(&choixClasse)

	switch choixClasse {
	case 1:
		classe = "Humain"
		pvMax = 100
	case 2:
		classe = "Elfe"
		pvMax = 80
	case 3:
		classe = "Nain"
		pvMax = 120
	default:
		fmt.Println("Choix invalide, vous serez Humain par défaut.")
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
