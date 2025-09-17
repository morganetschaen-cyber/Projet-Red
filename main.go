package main

import "fmt"

func main() {
	heros := characterCreation()

	for {
		fmt.Println("\n--- Menu principal ---")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l’inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Combat d'entraînement")
		fmt.Println("6. Quitter")

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
				fmt.Println("2. Utiliser une potion de poison")
				fmt.Println("3. Utiliser une augmentation d’inventaire")
				fmt.Println("4. Retour")

				var sub int
				fmt.Scan(&sub)

				if sub == 9 {
					break
				} else if sub == 1 {
					takePot(&heros)
				} else if sub == 2 {
					poisonPot(&heros)
					removeInventory(&heros, "Potion de poison")
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
			for gobelin.PVActuels > 0 && heros.PVActuels > 0 {
				characterTurn(&heros, &gobelin)
				if gobelin.PVActuels <= 0 {
					fmt.Println("Le gobelin est vaincu !")
					break
				}
				goblinAttack(&heros, &gobelin)
				if heros.PVActuels <= 0 {
					fmt.Println("Vous avez été vaincu !")
					break
				}
			}
		case 6:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
