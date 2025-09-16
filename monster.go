package main

import "fmt"

type Monster struct {
	Nom       string
	PVMax     int
	PVActuels int
	Attaque   int
}

func initGoblin() Monster {
	return Monster{
		Nom:       "Gobelin d’entraînement",
		PVMax:     40,
		PVActuels: 40,
		Attaque:   5,
	}
}
func goblinPattern(g Monster, c *Character, tour int) {
	degats := g.Attaque

	if tour%3 == 0 {
		degats = g.Attaque * 2
	}

	c.PVActuels -= degats
	if c.PVActuels < 0 {
		c.PVActuels = 0
	}

	fmt.Printf("%s inflige à %s %d de dégâts\n", g.Nom, c.Nom, degats)
	fmt.Printf("%s PV : %d/%d\n", c.Nom, c.PVActuels, c.PVMax)
}
