package main

import "fmt"

func marchand(c *Character) {
	for {
		fmt.Println("\n--- Marchand ---")
		fmt.Println("1. Potion de vie")
		fmt.Println("2. Retour au menu")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			addInventory(c, "Potion de vie")
		case 2:
			fmt.Println("Retour au menu principal.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
