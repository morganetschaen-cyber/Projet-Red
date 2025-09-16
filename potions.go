package main

import (
	"fmt"
	"time"
)

func poisonPot(c *Character) {
	for i := 1; i <= 3; i++ {
		c.PVActuels -= 10
		if c.PVActuels < 0 {
			c.PVActuels = 0
		}
		fmt.Printf("Poison! PV actuels: %d/%d\n", c.PVActuels, c.PVMax)
		time.Sleep(1 * time.Second)
	}
	isDead(c)
}
