package main

import "projet-red/game"

func main() {
	if !game.AfficherIntro() {
		return
	}

	heros := game.Creation()
	game.AfficherMenu(&heros)
}
