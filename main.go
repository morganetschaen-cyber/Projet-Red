package main

import "fmt"

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
			for {
				fmt.Println("\n--- Inventaire ---")
				accessInventory(heros.Inventaire)
				fmt.Println("1. Utiliser une potion")
				fmt.Println("9. Retour")

				var sub int
				fmt.Scan(&sub)

				if sub == 9 {
					break
				} else if sub == 1 {
					takePot(&heros)
				} else {
					fmt.Println("Choix invalide.")
				}
			}

		case 3:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
