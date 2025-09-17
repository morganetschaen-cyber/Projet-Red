package game

import (
	"fmt"
	"time"
)

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

func PoisonPot(c *Character, m *Monster) {
	for i, item := range c.Inventaire {
		if item == "Potion de poison" {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)

			fmt.Printf("%s lance une potion de poison sur %s !\n", c.Nom, m.Nom)

			for sec := 1; sec <= 3; sec++ {
				time.Sleep(1 * time.Second)
				m.PVActuels -= 10
				if m.PVActuels < 0 {
					m.PVActuels = 0
				}
				fmt.Printf("Seconde %d : %s a %d/%d PV\n", sec, m.Nom, m.PVActuels, m.PVMax)
			}
			return
		}
	}
	fmt.Println("Vous n’avez plus de potions de poison !")
}
