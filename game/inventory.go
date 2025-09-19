package game

import (
	"fmt"
	"strings"
)

func AccessInventory(inventory []string) {
	if len(inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	fmt.Println("Inventaire du personnage :")
	counts := make(map[string]int)
	for _, item := range inventory {
		counts[item]++
	}
	for item, qty := range counts {
		fmt.Printf("- %s (x%d)\n", item, qty)
	}
}

func AddInventory(c *Character, item string) {
	if CanAddItem(c) {
		c.Inventaire = append(c.Inventaire, item)
		fmt.Println(item, "a été ajouté à votre inventaire !")
	}
}

func RemoveInventory(c *Character, item string) {
	for i, v := range c.Inventaire {
		if v == item {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			fmt.Println(item, "a été retiré de l'inventaire.")
			return
		}
	}
	fmt.Println(item, "n'est pas dans l'inventaire.")
}

func CanAddItem(c *Character) bool {
	if len(c.Inventaire) >= c.CapaciteInventory {
		fmt.Println("Inventaire plein! Vous ne pouvez pas ajouter cet item.")
		return false
	}
	return true
}

func HasItem(inventory []string, item string) bool {
	for _, v := range inventory {
		if v == item {
			return true
		}
	}
	return false
}

const (
	reset = "\033[0m"

	gris  = "\033[90m"
	rouge = "\033[91m"
	jaune = "\033[93m"
	blanc = "\033[97m"
)

func AfficherInventaireCombat(c *Character, m *Monster) {
	for {
		width := 40
		top := "┌" + strings.Repeat("─", width-2) + "┐"
		mid := "├" + strings.Repeat("─", width-2) + "┤"
		bot := "└" + strings.Repeat("─", width-2) + "┘"

		fmt.Println(top)
		title := "INVENTAIRE DU VOYAGEUR"
		fmt.Printf("│ %-36s │\n", centerString(title, 36))
		fmt.Println(mid)

		if len(c.Inventaire) == 0 {
			fmt.Printf("│ %-36s │\n", centerString("(vide)", 36))
		} else {
			counts := make(map[string]int)
			for _, item := range c.Inventaire {
				counts[item]++
			}
			i := 1
			for item, qty := range counts {
				line := fmt.Sprintf("(%d) %-22s [x%d]", i, item, qty)
				fmt.Printf("│ %-36s │\n", line)
				i++
			}
		}

		fmt.Println(mid)
		goldLine := fmt.Sprintf("Argent : %d pièces", c.Argent)
		fmt.Printf("│ %-36s │\n", goldLine)
		fmt.Println(bot)

		fmt.Println()
		fmt.Println("> Que souhaitez-vous faire ?")
		if HasItem(c.Inventaire, "Potion") {
			fmt.Println("   (1) Utiliser une potion de vie")
		}
		if HasItem(c.Inventaire, "Potion de poison") {
			fmt.Println("   (2) Utiliser une potion de poison")
		}
		fmt.Println("   (0) Fermer l’inventaire")

		var sub int
		fmt.Print("> ")
		fmt.Scan(&sub)

		switch sub {
		case 1:
			if HasItem(c.Inventaire, "Potion") {
				TakePot(c)
			}
		case 2:
			if HasItem(c.Inventaire, "Potion de poison") {
				RemoveInventory(c, "Potion de poison")
			}
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func centerString(s string, width int) string {
	if len(s) >= width {
		return s
	}
	spaces := (width - len(s)) / 2
	return strings.Repeat(" ", spaces) + s
}
