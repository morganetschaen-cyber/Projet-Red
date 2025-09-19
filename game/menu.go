package game

import (
	"fmt"
	"strings"
)

func AfficherMenu(heros *Character) {
	width := 80

	for {
		fmt.Println()
		fmt.Println("╔" + strings.Repeat("═", width-2) + "╗")
		fmt.Println(centerText("✦ LES PORTES DE KRYNN ✦", width))
		fmt.Println("╚" + strings.Repeat("═", width-2) + "╝")
		fmt.Println()

		fmt.Println("Que désires-tu faire, voyageur ?")
		fmt.Println()

		fmt.Println("[1] Observer ton reflet")
		fmt.Println("[2] Examiner ton sac")
		fmt.Println("[3] Aller chez le marchand")
		fmt.Println("[4] Façonner ton équipement")
		fmt.Println("[5] Affronter tes démons")
		fmt.Println("[0] Quitter ce monde")

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
			fmt.Println("\nChoisis ton adversaire :")
			fmt.Println("[1] Gobelin")
			fmt.Println("[2] Troll")
			fmt.Println("[3] Dragon")
			var choixMonstre int
			fmt.Print("> ")
			fmt.Scan(&choixMonstre)

			var monstre Monster
			switch choixMonstre {
			case 1:
				monstre = InitGoblin()
			case 2:
				monstre = InitTroll()
			case 3:
				monstre = InitDragon()
			default:
				fmt.Println("Choix invalide.")
				continue
			}
			TrainingFight(heros, &monstre)

		case 0:
			fmt.Println("\nTu tournes le dos aux terres de Krynn...")
			fmt.Println("Le voyage s’achève ici.")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
