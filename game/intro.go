package game

import (
	"fmt"
	"strings"
	"time"
)

func slowPrint(s string) {
	for _, c := range s {
		fmt.Printf("%c", c)
		time.Sleep(30 * time.Millisecond)
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

func AfficherIntro() bool {
	width := 80

	logo := []string{
		"   ____  __.                            ",
		"  |    |/ _|______ ___.__. ____   ____  ",
		"  |      < \\_  __ <   |  |/    \\ /    \\ ",
		"  |    |  \\ |  | \\/\\___  |   |  \\   |  \\",
		"  |____|__ \\|__|   / ____|___|  /___|  /",
		"          \\/       \\/         \\/     \\/ ",
	}

	fmt.Println("\033[31m")
	fmt.Println("╔" + strings.Repeat("═", width-2) + "╗")
	for _, line := range logo {
		fmt.Printf("║%-78s║\n", centerText(line, 78))
	}
	fmt.Println("╚" + strings.Repeat("═", width-2) + "╝")
	fmt.Println("\033[0m")

	slowPrint("Bienvenue à Krynn.")
	slowPrint("Es-tu sûr de vouloir entrer ?")

	fmt.Println("\033[32m1. Oui\033[0m  |  \033[31m2. Non\033[0m")
	var choix int
	fmt.Print("> ")
	fmt.Scan(&choix)

	if choix == 2 {
		fmt.Println("\033[33m")
		slowPrint("Tu tournes le dos aux terres de Krynn...")
		fmt.Println("\033[0m")
		return false
	}

	fmt.Println("\033[31m")
	slowPrint("Ton voyage commence maintenant.")
	fmt.Println("\033[0m")
	return true
}
