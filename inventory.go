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
	for i, item := range c.Inventaire {
		if item == "Potion" {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			c.PVActuels += 50
			if c.PVActuels > c.PVMax {
				c.PVActuels = c.PVMax
			}
			fmt.Println("Vous utilisez une potion de soin !")
			fmt.Printf("PV : %d/%d\n", c.PVActuels, c.PVMax)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Vous n’avez plus de potions !")
	}
}

func addInventory(c *Character, item string) {
	if canAddItem(c) {
		c.Inventaire = append(c.Inventaire, item)
		fmt.Println(item, "a été ajouté à votre inventaire !")
	}
}

func removeInventory(c *Character, item string) {
	for i, v := range c.Inventaire {
		if v == item {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			fmt.Println(item, "a été retiré de l'inventaire.")
			return
		}
	}
	fmt.Println(item, "n'est pas dans l'inventaire.")
}

func canAddItem(c *Character) bool {
	if len(c.Inventaire) >= c.CapaciteInventaire {
		fmt.Println("Inventaire plein! Vous ne pouvez pas ajouter cet item.")
		return false
	}
	return true
}

func hasItem(inventory []string, item string) bool {
	for _, v := range inventory {
		if v == item {
			return true
		}
	}
	return false
}
