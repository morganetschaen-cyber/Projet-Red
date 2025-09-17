package game

import (
	"fmt"
	"math/rand"
)

func CharacterTurn(c *Character, m *Monster) {
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
		fmt.Printf("%s attaque et inflige 5 dégâts à %s.\n", c.Nom, m.Nom)
	case 2:
		AfficherInventaire(c)
	default:
		fmt.Println("Choix invalide, vous perdez votre tour !")
	}
}

func GoblinAttack(c *Character, m *Monster) {
	fmt.Println("\n--- Tour du monstre ---")
	c.PVActuels -= m.Attaque
	if c.PVActuels < 0 {
		c.PVActuels = 0
	}
	fmt.Printf("%s attaque %s et inflige %d dégâts.\n", m.Nom, c.Nom, m.Attaque)
	IsDead(c)
}

func TrainingFight(c *Character, m *Monster) {
	tour := 1
	for c.PVActuels > 0 && m.PVActuels > 0 {
		fmt.Printf("\n===== Tour %d =====\n", tour)

		CharacterTurn(c, m)
		if m.PVActuels <= 0 {
			fmt.Println(m.Nom, "est vaincu ! 🎉")
			goldEarned := rand.Intn(50) + 10
			c.Argent += goldEarned
			fmt.Printf("Vous recevez %d pièces d'or ! 💰\n", goldEarned)
			return
		}

		GoblinAttack(c, m)
		if IsDead(c) {
			return
		}
		tour++
	}
}
