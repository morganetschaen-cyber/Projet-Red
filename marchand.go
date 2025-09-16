package main

import "fmt"

func marchand(c *Character) {
	for {
		fmt.Println("\n--- Marchand ---")
		fmt.Println("1. Potion de vie")
		fmt.Println("2. Potion de poison")
		fmt.Println("3. Livre de Sort : Boule de Feu")
		fmt.Println("4. Retour au menu")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			addInventory(c, "Potion de vie")
		case 2:
			addInventory(c, "Potion de poison")
		case 3:
			addInventory(c, "Livre de sort : Boule de Feu")
			spellBook(c)
		case 4:
			fmt.Println("Retour au menu principal.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
