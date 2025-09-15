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

func initCharacter(nom, classe string, niveau, pvMax, pvActuels int, inventaire []string) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PVActuels:  pvActuels,
		Inventaire: inventaire,
	}
}

func displayInfo(c Character) {
	fmt.Println("Nom :", c.Nom)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Niveau)
	fmt.Println("Points de vie :", c.PVActuels, "/", c.PVMax)
	fmt.Println("Inventaire :", c.Inventaire)
}

func main() {
	inventaire := []string{"Potion", "Potion", "Potion"}
	heros := initCharacter("Eva", "Elfe", 1, 100, 50, inventaire)

	displayInfo(heros)
}
func main() {
	heros := initCharacter("TonNom", "Elfe", 1, 100, 40, []string{"Potion", "Potion", "Potion"})

	for {
		fmt.Println("\n--- Menu principal ---")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l’inventaire")
		fmt.Println("3. Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			displayInfo(heros)
		case 2:
		case 3:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
