package main

import "fmt"

func accessInventory(inventory map[string]int) {
	if len(inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	fmt.Println("Inventaire du personnage :")
	for item, qty := range inventory {
		fmt.Printf("- %s (x%d)\n", item, qty)
	}
}

func main() {
	inventory := map[string]int{}

	accessInventory(inventory)
}
