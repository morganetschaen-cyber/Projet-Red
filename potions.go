package main

import "fmt"

func takePot(c *Character) {
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

func poisonPot(c *Character, m *Monster) {
	for i, item := range c.Inventaire {
		if item == "Potion de poison" {

			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)

			m.PoisonTour = 3
			fmt.Printf("Vous lancez une potion de poison sur %s !\n", m.Nom)
			return
		}
	}
	fmt.Println("Vous n’avez plus de potions de poison !")
}
