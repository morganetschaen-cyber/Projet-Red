package main

import (
	"fmt"
	"strings"
	"time"
)

func slowPrint(s string) {
	for _, c := range s {
		fmt.Printf("%c", c)
		time.Sleep(40 * time.Millisecond)
	}
	fmt.Println()
}

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

	slowPrint("Bienvenue dans le monde de Krynn...")
	slowPrint("Êtes-vous prêt à entrer ?")

	fmt.Println("\033[32m1. Oui\033[0m  |  \033[31m2. Non\033[0m")
	var choix int
	fmt.Print("> ")
	fmt.Scan(&choix)

	if choix == 2 {
		slowPrint("Vous tournez le dos aux ténèbres... mais pour combien de temps ?")
		return false
	}

	slowPrint("...Très bien. Le destin vous attend.")
	return true
}
