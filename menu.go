package main

import "fmt"

func afficherMenu(heros *Character) {
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
			displayInfo(*heros)
			fmt.Println("\nAppuyez sur Entrée pour revenir au menu...")
			fmt.Scanln()
			fmt.Scanln()

		case 2:
			afficherInventaire(heros)

		case 3:
			marchand(heros)

		case 4:
			forgeron(heros)

		case 5:
			fmt.Println("Début du combat d’entraînement !")
			gobelin := initGoblin()
			trainingFight(heros, &gobelin)

		case 6:
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
