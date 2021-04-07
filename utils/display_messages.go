package utils

import (
	"fmt"
	"net/http"
)

// PrintText : Print any text on the terminal.
func PrintText(text string) {
	fmt.Println(text)
}

// DisplayGameRules : Displays the game rules on the browser.
func DisplayGameRules(writer http.ResponseWriter, request *http.Request) {
	gameRules := `
	GAME RULES !!!
	-----------------------------------
	1. Paper covers Rock --> PAPER WINS
	2. Rock smashes Scissor --> ROCK WINS
	3. Scissor cuts Paper --> SCISSOR WINS
	`
	fmt.Fprint(writer, gameRules)
	DisplaySelectPlayerOrComputer(writer, request)
}

// DisplaySelectPlayerOrComputer : Displays who to play with on the browser.
func DisplaySelectPlayerOrComputer(writer http.ResponseWriter, request *http.Request) {
	message := `
	Whom do you want to play with?

	Select 1 to play with another player
	Select 2 to play with computer
	`
	fmt.Fprint(writer, message)
}
