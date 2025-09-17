package main

import (
	"fmt"
	"math/rand"
)

var premierVisite bool = true

func marchand(c *Character) {
	for {
		fmt.Println("\n--- Marchand ---")
		fmt.Println("1. Potion de vie")
		fmt.Println("2. Potion de poison (6 pièces d'or)")
		fmt.Println("3. Livre de Sort : Boule de Feu (25 pièces d'or)")
		fmt.Println("4. Extension d'inventaire (20 pièces d'or)")
		fmt.Println("5. Vendre un objet en double")
		fmt.Println("6. Retour au menu")

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
			if c.Argent >= 20 {
				c.Argent -= 20
				c.CapaciteInventaire += 5
				fmt.Println("Votre inventaire peut contenir 5 objets de plus !")
			} else {
				fmt.Println("Pas assez d'or.")
			}
		case 5:
			sellDuplicates(c)
		case 6:
			fmt.Println("Vous quittez le magasin.")
			return

		default:
			fmt.Println("Choix invalide.")
		}
		fmt.Println("Argent actuel :", c.Argent)
	}
}
func sellDuplicates(c *Character) {
	counts := make(map[string]int)
	for _, item := range c.Inventaire {
		counts[item]++
	}
	duplicates := []string{}
	for item, qty := range counts {
		if qty > 1 {
			duplicates = append(duplicates, item)
		}
	}

	if len(duplicates) == 0 {
		fmt.Println("Vous n'avez pas de doublons à vendre.")
		return
	}

	fmt.Println("Objets en double que vous pouvez vendre :")
	for i, item := range duplicates {
		fmt.Printf("%d. %s (x%d)\n", i+1, item, counts[item])
	}

	var choice int
	fmt.Println("Choisissez l'objet à vendre :")
	fmt.Scan(&choice)

	if choice > 0 && choice <= len(duplicates) {
		selected := duplicates[choice-1]
		for i, v := range c.Inventaire {
			if v == selected {
				c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
				break
			}
		}
		gold := rand.Intn(20) + 5
		c.Argent += gold
		fmt.Println("Vous avez vendu 1", selected, "pour", gold, "pièces d'or ! 💰")
	} else {
		fmt.Println("Choix invalide.")
	}
}
