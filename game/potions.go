package game

import "fmt"

func TakePot(c *Character) {
	for i, item := range c.Inventaire {
		if item == "Potion" {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			heal := 20
			c.PVActuels += heal
			if c.PVActuels > c.PVMax {
				c.PVActuels = c.PVMax
			}
			fmt.Printf("Vous buvez une potion (+%d PV). PV : %d/%d\n", heal, c.PVActuels, c.PVMax)
			return
		}
	}
	fmt.Println("Vous n’avez plus de potions !")
}
