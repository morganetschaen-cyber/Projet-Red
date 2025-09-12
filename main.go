package main

import "fmt"

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PVActuels  int
	Inventaire []string
}

func main() {
	heros := Character{
		Nom:        "Eva",
		Classe:     "",
		Niveau:     1,
		PVMax:      100,
		PVActuels:  50,
		Inventaire: []string{"Potion", "Potion", "Potion"},
	}

	fmt.Println(heros)
}
