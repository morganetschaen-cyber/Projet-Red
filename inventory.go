package main

import "fmt"

func accessInventory(inventory []string) {
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

func takePot(c *Character) {
	found := false
	for _, item := range c.Inventaire {
		if item == "Potion" {
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Vous n’avez plus de potions !")
	}
}
