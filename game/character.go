package game

import (
	"fmt"
	"strings"
)

type Character struct {
	Nom               string
	Classe            string
	Niveau            int
	PVMax             int
	PVActuels         int
	Inventaire        []string
	Skill             []string
	Argent            int
	Equipment         Equipment
	CapaciteInventory int
}

func center(text string, width int) string {
	if len(text) >= width {
		return text
	}
	spaces := (width - len(text)) / 2
	return strings.Repeat(" ", spaces) + text
}

func Creation() Character {
	width := 80
	fmt.Println(strings.Repeat("=", width))
	fmt.Println(center("✦ CRÉATION DU VOYAGEUR ✦", width))
	fmt.Println(strings.Repeat("=", width))

	var nom string
	fmt.Println("\nEntrez le nom du voyageur :")
	fmt.Print("> ")
	fmt.Scan(&nom)
	if len(nom) > 0 {
		nom = strings.ToUpper(nom[:1]) + strings.ToLower(nom[1:])
	}

	for {
		fmt.Println("\nChoisis ton chemin :")
		fmt.Println("  [1] Assassin      → Rapide, précis, impitoyable")
		fmt.Println("  [2] Écorcheur     → Sauvage, brutal, sanguinaire")
		fmt.Println("  [3] Nécromancien  → Sombre, mystique, impie")
		fmt.Print("> ")

		var choix int
		fmt.Scan(&choix)

		var classe string
		var pvMax int
		switch choix {
		case 1:
			classe, pvMax = "Assassin", 90
		case 2:
			classe, pvMax = "Écorcheur", 110
		case 3:
			classe, pvMax = "Nécromancien", 80
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		fmt.Printf("\nTu as choisi : %s\n", classe)
		fmt.Println("Es-tu sûr ? (1. Oui | 2. Non)")
		fmt.Print("> ")

		var confirm int
		fmt.Scan(&confirm)

		if confirm == 1 {
			fmt.Println()
			fmt.Println(center("✦ Tu es maintenant liée aux secrets de Krynn. ✦", width))

			return Character{
				Nom:               nom,
				Classe:            classe,
				Niveau:            1,
				PVMax:             pvMax,
				PVActuels:         pvMax / 2,
				Inventaire:        []string{},
				Skill:             []string{"Coup de poing"},
				Argent:            100,
				Equipment:         Equipment{},
				CapaciteInventory: 10,
			}
		}
	}
}

func DisplayInfo(c Character) {
	width := 60
	fmt.Println("╔" + strings.Repeat("═", width-2) + "╗")
	fmt.Println(center("✦ TON REFLET ✦", width))
	fmt.Println("╚" + strings.Repeat("═", width-2) + "╝")

	fmt.Printf("Nom       : %s\n", c.Nom)
	fmt.Printf("Classe    : %s\n", c.Classe)
	fmt.Printf("Niveau    : %d\n", c.Niveau)
	fmt.Printf("PV        : %d / %d\n", c.PVActuels, c.PVMax)
	fmt.Printf("Inventaire: %v\n", c.Inventaire)
	fmt.Printf("Argent    : %d pièces\n", c.Argent)
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

func RecalculerPVMax(c *Character) {
	var basePV int
	switch strings.ToLower(c.Classe) {
	case "assassin":
		basePV = 90
	case "écorcheur":
		basePV = 110
	case "nécromancien":
		basePV = 80
	default:
		basePV = 100
	}
	basePV += (c.Niveau - 1) * 10

	if c.Equipment.Tete == "Chapeau de l’aventurier" {
		basePV += 10
	}
	if c.Equipment.Torse == "Tunique de l’aventurier" {
		basePV += 25
	}
	if c.Equipment.Pieds == "Bottes de l’aventurier" {
		basePV += 15
	}

	c.PVMax = basePV
	if c.PVActuels > c.PVMax {
		c.PVActuels = c.PVMax
	}
}
