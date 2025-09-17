package main

import (
	"fmt"
	"math/rand"
)

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
			fmt.Println("2. Lancer une potion de poison sur (l'ennemi)")
			fmt.Println("9. Retour")

			var choixInv int
			fmt.Scan(&choixInv)

			if choixInv == 9 {
				break
			} else if choixInv == 1 {
				takePot(c)
				break
			} else if choixInv == 2 {
				poisonPot(c, m)
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

	if m.PoisonTour > 0 {
		damage := 10
		m.PVActuels -= damage
		if m.PVActuels < 0 {
			m.PVActuels = 0
		}
		fmt.Printf("💀 Le poison ronge %s ! (-%d PV, reste %d/%d)\n", m.Nom, damage, m.PVActuels, m.PVMax)
		m.PoisonTour--

		if m.PVActuels == 0 {
			fmt.Printf("%s succombe au poison...\n", m.Nom)
			return
		}
	}

	c.PVActuels -= m.Attaque
	if c.PVActuels < 0 {
		c.PVActuels = 0
	}
	fmt.Printf("%s attaque %s et inflige %d dégâts.\n", m.Nom, c.Nom, m.Attaque)
	fmt.Printf("%s PV: %d/%d\n", c.Nom, c.PVActuels, c.PVMax)
	isDead(c)
}
func trainingFight(c *Character, m *Monster) {
	tour := 1
	for c.PVActuels > 0 && m.PVActuels > 0 {
		fmt.Printf("\n===== Tour %d =====\n", tour)

		characterTurn(c, m)
		if m.PVActuels <= 0 {
			fmt.Println(m.Nom, "est vaincu ! 🎉")
			goldEarned := rand.Intn(50) + 10
			c.Argent += goldEarned
			fmt.Printf("Vous recevez %d pièces d'or ! 💰\n", goldEarned)
			return
		}

		goblinAttack(c, m)
		if isDead(c) {
			return
		}

		tour++
	}
}
