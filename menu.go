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
			for {
				fmt.Println("\n--- Inventaire ---")
				accessInventory(heros.Inventaire)

				options := make(map[int]string)
				optionNum := 1

				if hasItem(heros.Inventaire, "Potion") {
					options[optionNum] = "Utiliser une potion de vie"
					optionNum++
				}

				equipables := []string{"Chapeau de l’aventurier", "Tunique de l’aventurier", "Bottes de l’aventurier"}
				for _, item := range equipables {
					if hasItem(heros.Inventaire, item) {
						options[optionNum] = "Équiper " + item
						optionNum++
					}
				}

				options[optionNum] = "Retour"

				for k, v := range options {
					fmt.Printf("%d. %s\n", k, v)
				}

				var sub int
				fmt.Print("> ")
				fmt.Scan(&sub)

				action, exists := options[sub]
				if !exists {
					fmt.Println("Choix invalide.")
					continue
				}

				switch action {
				case "Retour":
					break
				case "Utiliser une potion de vie":
					takePot(heros)
				default:
					equipItem(heros, action[len("Équiper "):])
				}

				if action == "Retour" {
					break
				}
			}

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
