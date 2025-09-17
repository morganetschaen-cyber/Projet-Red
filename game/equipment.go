package game

import "fmt"

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

func EquipItem(c *Character, item string) {
	switch item {
	case "Chapeau de l’aventurier":
		if c.Equipment.Tete != "" {
			AddInventory(c, c.Equipment.Tete)
		}
		c.Equipment.Tete = item
		c.PVMax += 10
	case "Tunique de l’aventurier":
		if c.Equipment.Torse != "" {
			AddInventory(c, c.Equipment.Torse)
		}
		c.Equipment.Torse = item
		c.PVMax += 25
	case "Bottes de l’aventurier":
		if c.Equipment.Pieds != "" {
			AddInventory(c, c.Equipment.Pieds)
		}
		c.Equipment.Pieds = item
		c.PVMax += 15
	default:
		fmt.Println("Cet objet ne peut pas être équipé.")
		return
	}

	RemoveInventory(c, item)
	fmt.Println(item, "équipé avec succès !")
	fmt.Printf("PV Max: %d\n", c.PVMax)
}
