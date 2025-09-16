package main

import "fmt"

var premierVisite bool = true

func marchand(c *Character) {
	for {
		fmt.Println("\n--- Marchand ---")
		fmt.Println("1. Potion de vie")
		fmt.Println("2. Potion de poison (6 pièces d'or)")
		fmt.Println("3. Livre de Sort : Boule de Feu (25 pièces d'or)")
		fmt.Println("4. Retour au menu")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if premierVisite {
				fmt.Println("Cette potion est gratuite pour votre première visite !")
				addInventory(c, "Potion de vie")
				premierVisite = false
			} else if c.Argent >= 3 {
				c.Argent -= 3
				addInventory(c, "Potion de vie")
			} else {
				fmt.Println("Pas assez d'argent !")
			}
		case 2:
			if c.Argent >= 6 {
				c.Argent -= 6
				addInventory(c, "Potion de poison")
			} else {
				fmt.Println("Pas assez d'argent !")
			}
		case 3:
			if c.Argent >= 25 {
				c.Argent -= 25
				addInventory(c, "Livre de Sort: Boule de feu")
				spellBook(c)
			} else {
				fmt.Println("Pas assez d'argent !")
			}
		case 4:
			fmt.Println("Retour au menu principal.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
		fmt.Println("Argent actuel :", c.Argent)
	}
}
