package game

import "fmt"

func AfficherMenu(heros *Character) {
	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l’inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Combat d’entraînement")
		fmt.Println("0. Quitter")

		var choix int
		fmt.Print("> ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			DisplayInfo(*heros)
		case 2:
			AfficherInventaire(heros)
		case 3:
			Marchand(heros)
		case 4:
			Forgeron(heros)
		case 5:
			gobelin := InitGoblin()
			TrainingFight(heros, &gobelin)
		case 6:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
