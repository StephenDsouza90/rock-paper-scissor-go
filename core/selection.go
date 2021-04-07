package core

import (
	"math/rand"
)

// GetComputerSelection : Computer's selection taken from the game option dict.
func GetComputerSelection() string {
	randomIndex := rand.Intn(len(GameOptions))
	computerPick := GameOptions[randomIndex]
	return computerPick
}

// GetPlayerSelection : Player's selection taken from the game option dict.
func GetPlayerSelection(playerInput int) string {
	playerPick := GameOptions[playerInput]
	return playerPick
}
