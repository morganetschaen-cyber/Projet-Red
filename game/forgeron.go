package game

import "fmt"

type Recipe struct {
	Nom        string
	Ressources []string
	Prix       int
}

var recipes = []Recipe{
	{"Chapeau de l’aventurier", []string{"Plume de Corbeau", "Cuir de Sanglier"}, 5},
	{"Tunique de l’aventurier", []string{"Fourrure de Loup", "Fourrure de Loup", "Peau de Troll"}, 5},
	{"Bottes de l’aventurier", []string{"Fourrure de Loup", "Cuir de Sanglier"}, 5},
}

func Forgeron(c *Character) {
	for {
		fmt.Println("\n--- Forgeron ---")
		for i, r := range recipes {
			fmt.Printf("%d. %s (", i+1, r.Nom)
			for j, res := range r.Ressources {
				if j > 0 {
					fmt.Print(", ")
				}
				fmt.Print(res)
			}
			fmt.Printf(", %d or)\n", r.Prix)
		}
		fmt.Println("0. Retour")

		var choix int
		fmt.Scan(&choix)

		if choix == 0 {
			return
		}
		if choix > 0 && choix <= len(recipes) {
			r := recipes[choix-1]
			if craftItem(c, r.Nom, r.Ressources, r.Prix) {
				fmt.Println("Vous avez fabriqué :", r.Nom)
			}
		} else {
			fmt.Println("Choix invalide.")
		}
	}
}

func craftItem(c *Character, item string, ressources []string, prix int) bool {
	if c.Argent < prix {
		fmt.Println("Pas assez d’or.")
		return false
	}

	for _, r := range ressources {
		if !HasItem(c.Inventaire, r) {
			fmt.Println("Ressource manquante :", r)
			return false
		}
	}

	for _, r := range ressources {
		RemoveInventory(c, r)
	}
	c.Argent -= prix
	AddInventory(c, item)

	return true
}
