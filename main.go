package main

func main() {
	if !afficherIntro() {
		return
	}

	heros := characterCreation()
	afficherMenu(&heros)
}
