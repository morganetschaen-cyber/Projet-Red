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
		Niveau:     0,
		PVMax:      0,
		PVActuels:  0,
		Inventaire: []string{},
	}

	fmt.Println(heros)
}
