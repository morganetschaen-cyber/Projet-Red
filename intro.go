package main

import (
	"fmt"
	"strings"
)

func centerText(text string, width int) string {
	if len(text) >= width {
		return text
	}
	spaces := (width - len(text)) / 2
	return strings.Repeat(" ", spaces) + text
}

func afficherIntro() bool {
	width := 80
	fmt.Println("\033[31m")
	fmt.Println(centerText("+--------------------------------------------------+", width))
	fmt.Println(centerText("!     ____  __.                            ", width))
	fmt.Println(centerText("!    |    |/ _|______ ___.__. ____   ____  ", width))
	fmt.Println(centerText("!    |      < \\_  __ <   |  |/    \\ /    \\ ", width))
	fmt.Println(centerText("!    |    |  \\ |  | \\/\\___  |   |  \\   |  \\", width))
	fmt.Println(centerText("!    |____|__ \\|__|   / ____|___|  /___|  /", width))
	fmt.Println(centerText("!            \\/       \\/         \\/     \\/ ", width))
	fmt.Println(centerText("+--------------------------------------------------+", width))
	fmt.Println("\033[0m")
	return true

	fmt.Println("Bienvenue dans le monde de Krynn...")
	fmt.Println("Êtes-vous prêt à entrer ?")
	fmt.Println("1. Oui  |  2. Non")

	var choix int
	fmt.Print("> ")
	fmt.Scan(&choix)

	if choix == 2 {
		fmt.Println("Vous tournez le dos aux ténèbres... mais pour combien de temps ?")
		return false
	}

	fmt.Println("...Très bien. Le destin vous attend.")
	return true
}
