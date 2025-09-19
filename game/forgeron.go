package game

import (
	"fmt"
	"strings"
)

func Forgeron(c *Character) {
	for {
		width := 50
		top := "┌" + strings.Repeat("─", width-2) + "┐"
		sep := "├" + strings.Repeat("─", width-2) + "┤"
		bot := "└" + strings.Repeat("─", width-2) + "┘"

		fmt.Println(top)
		fmt.Printf("│ %-46s │\n", center("⚒ FORGERON ⚒", 46))
		fmt.Println(sep)

		fmt.Printf("│ %-46s │\n", "1. Chapeau de l’aventurier")
		fmt.Printf("│ %-46s │\n", "   → Plume de Corbeau, Cuir de Sanglier (5 or)")
		fmt.Printf("│ %-46s │\n", "")

		fmt.Printf("│ %-46s │\n", "2. Tunique de l’aventurier")
		fmt.Printf("│ %-46s │\n", "   → Fourrure de Loup x2, Peau de Troll (5 or)")
		fmt.Printf("│ %-46s │\n", "")

		fmt.Printf("│ %-46s │\n", "3. Bottes de l’aventurier")
		fmt.Printf("│ %-46s │\n", "   → Fourrure de Loup, Cuir de Sanglier (5 or)")
		fmt.Printf("│ %-46s │\n", "")

		fmt.Printf("│ %-46s │\n", "0. Retour")

		fmt.Println(bot)

		var choix int
		fmt.Print("> ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if craftItem(c, "Chapeau de l’aventurier",
				[]string{"Plume de Corbeau", "Cuir de Sanglier"}, 5) {
				fmt.Println("Vous avez fabriqué un Chapeau de l’aventurier !")
			}
		case 2:
			if craftItem(c, "Tunique de l’aventurier",
				[]string{"Fourrure de Loup", "Fourrure de Loup", "Peau de Troll"}, 5) {
				fmt.Println("Vous avez fabriqué une Tunique de l’aventurier !")
			}
		case 3:
			if craftItem(c, "Bottes de l’aventurier",
				[]string{"Fourrure de Loup", "Cuir de Sanglier"}, 5) {
				fmt.Println("Vous avez fabriqué des Bottes de l’aventurier !")
			}
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func craftItem(c *Character, item string, ressources []string, prix int) bool {
	if c.Argent < prix {
		fmt.Println("Pas assez d’or.")
		return false
	}

	for _, r := range ressources {
		if !HasItem(c.Inventaire, r) {
			fmt.Println("Ressource manquante :", r)
			return false
		}
	}

	for _, r := range ressources {
		RemoveInventory(c, r)
	}

	c.Argent -= prix
	AddInventory(c, item)

	return true
}
