package main

import "fmt"

func accessInventory(inventory []string) {
	if len(inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}
}
