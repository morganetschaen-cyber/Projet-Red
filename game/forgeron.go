package game

import (
	"fmt"
	"strings"
)

func Forgeron(c *Character) {
	width := 80

	for {
		fmt.Println()
		fmt.Println("╔" + strings.Repeat("═", width-2) + "╗")
		fmt.Println(center("✦ FORGERON ✦", width))
		fmt.Println("╚" + strings.Repeat("═", width-2) + "╝")

		fmt.Println("Que veux-tu façonner ?")
		fmt.Println()
		fmt.Println("[1] Chapeau de l’aventurier  → Plume de Corbeau, Cuir de Sanglier")
		fmt.Println("[2] Tunique de l’aventurier  → Fourrure de Loup x2, Peau de Troll")
		fmt.Println("[3] Bottes de l’aventurier   → Fourrure de Loup, Cuir de Sanglier")
		fmt.Println("[0] Retour")

		var choix int
		fmt.Print("> ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if craftItem(c, "Chapeau de l’aventurier",
				[]string{"Plume de Corbeau", "Cuir de Sanglier"}, 5) {
				fmt.Println("Vous avez façonné un Chapeau de l’aventurier !")
				fmt.Println("Voulez-vous l’équiper ? (1=Oui / 0=Non)")
				var rep int
				fmt.Scan(&rep)
				if rep == 1 {
					EquipItem(c, "Chapeau de l’aventurier")
				}
			}
		case 2:
			if craftItem(c, "Tunique de l’aventurier",
				[]string{"Fourrure de Loup", "Fourrure de Loup", "Peau de Troll"}, 5) {
				fmt.Println("Vous avez façonné une Tunique de l’aventurier !")
				fmt.Println("Voulez-vous l’équiper ? (1=Oui / 0=Non)")
				var rep int
				fmt.Scan(&rep)
				if rep == 1 {
					EquipItem(c, "Tunique de l’aventurier")
				}
			}
		case 3:
			if craftItem(c, "Bottes de l’aventurier",
				[]string{"Fourrure de Loup", "Cuir de Sanglier"}, 5) {
				fmt.Println("Vous avez façonné des Bottes de l’aventurier !")
				fmt.Println("Voulez-vous l’équiper ? (1=Oui / 0=Non)")
				var rep int
				fmt.Scan(&rep)
				if rep == 1 {
					EquipItem(c, "Bottes de l’aventurier")
				}
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

