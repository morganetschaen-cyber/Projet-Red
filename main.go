package main

import "fmt"

func main() {
	heros := characterCreation()

	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l’inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Combat d’entraînement")
		fmt.Println("6. Quitter")

		var choix int
		fmt.Print("> ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			displayInfo(heros)

			fmt.Println("\nAppuyez sur Entrée pour revenir au menu...")
			fmt.Scanln()
			fmt.Scanln()

		case 2:
			for {
				fmt.Println("\n--- Inventaire ---")
				accessInventory(heros.Inventaire)

				if hasItem(heros.Inventaire, "Potion") {
					fmt.Println("1. Utiliser une potion de vie")
				}
				fmt.Println("2. Retour")

				var sub int
				fmt.Scan(&sub)

				if sub == 2 {
					break
				} else if sub == 1 && hasItem(heros.Inventaire, "Potion") {
					takePot(&heros)
				} else {
					fmt.Println("Choix invalide.")
				}
			}

		case 3:
			marchand(&heros)

		case 4:
			forgeron(&heros)

		case 5:
			fmt.Println("Début du combat d’entraînement !")
			gobelin := initGoblin()
			trainingFight(&heros, &gobelin)

		case 6:
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
