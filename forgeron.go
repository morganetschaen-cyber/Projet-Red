package main

import "fmt"

func forgeron(c *Character) {
	for {
		fmt.Println("\n--- Forgeron ---")
		fmt.Println("1. Chapeau de l’aventurier (5 pièces d'or)")
		fmt.Println("2. Tunique de l’aventurier (5 pièces d'or)")
		fmt.Println("3. Bottes de l’aventurier (5 pièces d'or)")
		fmt.Println("4. Retour au menu")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if c.Argent >= 5 {
				c.Argent -= 5
				addInventory(c, "Chapeau de l’aventurier")
			} else {
				fmt.Println("Pas assez d'argent !")
			}
		case 2:
			if c.Argent >= 5 {
				c.Argent -= 5
				addInventory(c, "Tunique de l’aventurier")
			} else {
				fmt.Println("Pas assez d'argent !")
			}
		case 3:
			if c.Argent >= 5 {
				c.Argent -= 5
				addInventory(c, "Bottes de l’aventurier")
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
