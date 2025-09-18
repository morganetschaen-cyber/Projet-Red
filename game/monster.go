package game

import "fmt"

type Monster struct {
	Nom        string
	PVMax      int
	PVActuels  int
	Attaque    int
	PoisonTour int
}

// === INIT MONSTRES ===
func InitGoblin() Monster {
	return Monster{
		Nom:       "Gobelin d’entraînement",
		PVMax:     40,
		PVActuels: 40,
		Attaque:   5,
	}
}

func InitTroll() Monster {
	return Monster{
		Nom:       "Troll des cavernes",
		PVMax:     80,
		PVActuels: 80,
		Attaque:   12,
	}
}

func InitDragon() Monster {
	return Monster{
		Nom:       "Dragon de feu",
		PVMax:     200,
		PVActuels: 200,
		Attaque:   25,
	}
}

// === PATTERN GENERAL ===
func MonsterPattern(m *Monster, c *Character, tour int) {
	switch m.Nom {
	case "Gobelin d’entraînement":
		dmg := m.Attaque
		if tour%3 == 0 {
			dmg *= 2
			fmt.Println(m.Nom, "fait une attaque CRITIQUE !")
		}
		c.PVActuels -= dmg
		fmt.Printf("%s inflige %d dégâts à %s\n", m.Nom, dmg, c.Nom)

	case "Troll des cavernes":
		if tour%2 == 0 {
			fmt.Println(m.Nom, "se repose et ne fait rien ce tour...")
			return
		}
		dmg := m.Attaque * 2
		c.PVActuels -= dmg
		fmt.Printf("%s écrase %s et inflige %d dégâts\n", m.Nom, c.Nom, dmg)

	case "Dragon de feu":
		dmg := m.Attaque
		c.PVActuels -= dmg
		fmt.Printf("%s griffe %s et inflige %d dégâts\n", m.Nom, c.Nom, dmg)

		if tour%5 == 0 {
			extra := 15
			c.PVActuels -= extra
			fmt.Printf("🔥 %s crache du feu et inflige %d dégâts supplémentaires !\n",
				m.Nom, extra)
		}

	default:
		dmg := m.Attaque
		c.PVActuels -= dmg
		fmt.Printf("%s attaque %s et inflige %d dégâts\n", m.Nom, c.Nom, dmg)
	}

	if c.PVActuels < 0 {
		c.PVActuels = 0
	}
	fmt.Printf("%s PV : %d/%d\n", c.Nom, c.PVActuels, c.PVMax)
}
