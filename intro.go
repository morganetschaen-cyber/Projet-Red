package main

import "fmt"

func afficherIntro() bool {
	fmt.Println("\033[31m"
	fmt.Println("!     ____  __.                            ")
	fmt.Println("!    |    |/ _|______ ___.__. ____   ____  ")
	fmt.Println("!    |      < \\_  __ <   |  |/    \\ /    \\ ")
	fmt.Println("!    |    |  \\ |  | \\/\\___  |   |  \\   |  \\")
	fmt.Println("!    |____|__ \\|__|   / ____|___|  /___|  /")
	fmt.Println("!            \\/       \\/         \\/     \\/ ")
	fmt.Println("\033[0m")
	return true
}
