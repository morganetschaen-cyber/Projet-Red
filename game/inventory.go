package game

import (
	"fmt"
)

func AccessInventory(inventory []string) {
	if len(inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	fmt.Println("Inventaire du personnage :")
	counts := make(map[string]int)
	for _, item := range inventory {
		counts[item]++
	}
	for item, qty := range counts {
		fmt.Printf("- %s (x%d)\n", item, qty)
	}
}

func AddInventory(c *Character, item string) {
	if CanAddItem(c) {
		c.Inventaire = append(c.Inventaire, item)
		fmt.Println(item, "a été ajouté à votre inventaire !")
	}
}

func RemoveInventory(c *Character, item string) {
	for i, v := range c.Inventaire {
		if v == item {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			fmt.Println(item, "a été retiré de l'inventaire.")
			return
		}
	}
	fmt.Println(item, "n'est pas dans l'inventaire.")
}

func CanAddItem(c *Character) bool {
	if len(c.Inventaire) >= c.CapaciteInventaire {
		fmt.Println("Inventaire plein! Vous ne pouvez pas ajouter cet item.")
		return false
	}
	return true
}

func HasItem(inventory []string, item string) bool {
	for _, v := range inventory {
		if v == item {
			return true
		}
	}
	return false
}

func AfficherInventaire(c *Character) {
	for {
		fmt.Println("\n--- Inventaire ---")
		AccessInventory(c.Inventaire)

		if HasItem(c.Inventaire, "Potion") {
			fmt.Println("1. Utiliser une potion de vie")
		}
		if HasItem(c.Inventaire, "Potion de poison") {
			fmt.Println("2. Utiliser une potion de poison")
		}
		fmt.Println("0. Retour")

		var sub int
		fmt.Print("> ")
		fmt.Scan(&sub)

		switch sub {
		case 1:
			if HasItem(c.Inventaire, "Potion") {
				TakePot(c)
			}
		case 2:
			if HasItem(c.Inventaire, "Potion de poison") {
				RemoveInventory(c, "Potion de poison")
			}
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
