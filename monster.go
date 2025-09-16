package main

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
}
