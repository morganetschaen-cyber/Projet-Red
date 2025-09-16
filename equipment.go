package main

import "fmt"

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

func equipItem(c *Character, item string) {
	switch item {
	case "Chapeau de l’aventurier":
		if c.Equipment.Tete != "" {
			addInventory(c, c.Equipment.Tete) // remet l'ancien dans l'inventaire
		}
		c.Equipment.Tete = item
		c.PVMax += 10
	case "Tunique de l’aventurier":
		if c.Equipment.Torse != "" {
			addInventory(c, c.Equipment.Torse)
		}
		c.Equipment.Torse = item
		c.PVMax += 25
	case "Bottes de l’aventurier":
		if c.Equipment.Pieds != "" {
			addInventory(c, c.Equipment.Pieds)
		}
		c.Equipment.Pieds = item
		c.PVMax += 15
	default:
		fmt.Println("Cet objet ne peut pas être équipé.")
		return
	}

	removeInventory(c, item)
	fmt.Println(item, "équipé avec succès !")
	fmt.Printf("PV Max: %d\n", c.PVMax)
}
