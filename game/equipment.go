package game

import "fmt"

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

func EquipItem(c *Character, item string) {
	found := false
	for _, v := range c.Inventaire {
		if v == item {
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Vous n'avez pas cet objet dans votre inventaire.")
		return
	}

	switch item {
	case "Chapeau de l’aventurier":
		if c.Equipment.Tete != "" {
			c.Inventaire = append(c.Inventaire, c.Equipment.Tete)
		}
		c.Equipment.Tete = item

	case "Tunique de l’aventurier":
		if c.Equipment.Torse != "" {
			c.Inventaire = append(c.Inventaire, c.Equipment.Torse)
		}
		c.Equipment.Torse = item

	case "Bottes de l’aventurier":
		if c.Equipment.Pieds != "" {
			c.Inventaire = append(c.Inventaire, c.Equipment.Pieds)
		}
		c.Equipment.Pieds = item

	default:
		fmt.Println("Cet objet ne peut pas être équipé.")
		return
	}

	RemoveInventory(c, item)
	RecalculerPVMax(c)

	fmt.Println(item, "équipé avec succès !")
	fmt.Printf("PV Max: %d\n", c.PVMax)
}
