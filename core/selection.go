package core

import (
	"math/rand"
)

func ComputerSelection() string {
	/* FIXME Sometimes it is blank */
	/* rand.Seed(time.Now().UnixNano())  https://golang.cafe/blog/golang-random-number-generator.html */

	randomIndex := rand.Intn(len(GameOptions))
	computerPick := GameOptions[randomIndex]
	return computerPick
}

func PlayerSelection(playerInput int) string {
	playerPick := GameOptions[playerInput]
	return playerPick
}
