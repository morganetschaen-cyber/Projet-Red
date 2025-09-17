package main

import "fmt"

func characterTurn(c *Character, m *Monster) {
	fmt.Println("\n--- Tour du joueur ---")
	fmt.Println("1. Attaquer (5 dégâts)")
	fmt.Println("2. Inventaire")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		m.PVActuels -= 5
		if m.PVActuels < 0 {
			m.PVActuels = 0
		}
		fmt.Printf("%s utilise Attaque basique et inflige 5 dégâts à %s.\n", c.Nom, m.Nom)
		fmt.Printf("%s PV: %d/%d\n", m.Nom, m.PVActuels, m.PVMax)

	case 2:
		for {
			fmt.Println("\n--- Inventaire ---")
			accessInventory(c.Inventaire)
			fmt.Println("1. Utiliser une potion de vie")
			fmt.Println("2. Utiliser une potion de poison (sur vous)")
			fmt.Println("9. Retour")

			var choixInv int
			fmt.Scan(&choixInv)

			if choixInv == 9 {
				break
			} else if choixInv == 1 {
				takePot(c)
				break
			} else if choixInv == 2 {
				poisonPot(c)
				removeInventory(c, "Potion de poison")
				break
			} else {
				fmt.Println("Choix invalide.")
			}
		}

	default:
		fmt.Println("Choix invalide, vous perdez votre tour !")
	}
}

func goblinAttack(c *Character, m *Monster) {
	fmt.Println("\n--- Tour du monstre ---")
	c.PVActuels -= m.Attaque
	if c.PVActuels < 0 {
		c.PVActuels = 0
	}
	fmt.Printf("%s attaque %s et inflige %d dégâts.\n", m.Nom, c.Nom, m.Attaque)
	fmt.Printf("%s PV: %d/%d\n", c.Nom, c.PVActuels, c.PVMax)
	isDead(c)
}
