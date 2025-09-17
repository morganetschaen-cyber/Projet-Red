package game

import "fmt"

func Forgeron(c *Character) {
	for {
		fmt.Println("\n--- Forgeron ---")
		fmt.Println("1. Chapeau de l’aventurier")
		fmt.Println("2. Tunique de l’aventurier")
		fmt.Println("3. Bottes de l’aventurier")
		fmt.Println("4. Retour")

		var choix int
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
		case 4:
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
