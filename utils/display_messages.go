package utils

import (
	"fmt"
	"net/http"
)

func PrintText(text string) {
	fmt.Println(text)
}

func DisplayGameRules(writer http.ResponseWriter, request *http.Request) {
	gameRules := `
	GAME RULES !!!
	-----------------------------------
	1. Paper covers Rock --> PAPER WINS
	2. Rock smashes Scissor --> ROCK WINS
	3. Scissor cuts Paper --> Siccor WINS
	`
	fmt.Fprint(writer, gameRules)
	PrintText("Endpoint Hit: DisplayGameRules")
	/* PrintText(gameRules) */

	DisplaySelectPlayerOrComputer(writer, request)
}

func DisplaySelectPlayerOrComputer(writer http.ResponseWriter, request *http.Request) {
	message := `
	Whom do you want to play with?

	Select 1 to play with another player
	Select 2 to play with computer
	`
	fmt.Fprint(writer, message)
	PrintText("Endpoint Hit: DisplaySelectPlayerOrComputer")
	/* PrintText(message) */
}

func DisplayExitMessage() {
	exitMessage := "Game exited!"
	PrintText(exitMessage)
}

func DisplayUserOptionsMessage() {
	message := `
	Select one of the following:

	Press 1 for Rock
	Press 2 for Paper
	Press 3 for Scissor
	Press 0 to exit
	`
	PrintText(message)
}
