package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/StephenDsouza90/RockPaperScissor/core"
	"github.com/StephenDsouza90/RockPaperScissor/models"
	"github.com/StephenDsouza90/RockPaperScissor/utils"
)

/*
TODO:
1. Count scores
2. Build front-end
*/

func init() {
	/* https://stackoverflow.com/questions/20508356/non-declaration-statement-outside-function-body-in-go
	Alternatively, for more complex package initializations or
	to set up whatever state is required by the package GO provides an init function.
	Init will be called before main is run. */

	core.GameOptions = make(map[int]string)
	core.GameOptions[1] = "Rock"
	core.GameOptions[2] = "Paper"
	core.GameOptions[3] = "Scissor"
}

func selectPlayerOrComputer(writer http.ResponseWriter, request *http.Request) {
	// curl -X POST "http://127.0.0.1:8080/select-player-or-computer" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"PlayWith\": 1}"
	// curl -X POST "http://127.0.0.1:8080/select-player-or-computer" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"PlayWith\": 2}"

	var PlayWith models.PlayWithOptions

	const playWithPlayer = 1
	const playWithComputer = 2

	requestBody, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(requestBody, &PlayWith)
	utils.HandleError(err)

	playWith := PlayWith.PlayWith
	if playWith == playWithPlayer {
		response := models.Message{Message: "You are now playing with another player"}
		json.NewEncoder(writer).Encode(response)
		utils.DisplayUserOptionsMessage()
	} else if playWith == playWithComputer {
		response := models.Message{Message: "You are now playing with the computer"}
		json.NewEncoder(writer).Encode(response)
		utils.DisplayUserOptionsMessage()
	}
}

func startGameWithComputer(writer http.ResponseWriter, request *http.Request) {
	// curl -X POST "http://127.0.0.1:8080/start-game-with-computer" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"PlayerInput\": 1}"
	// curl -X POST "http://127.0.0.1:8080/start-game-with-computer" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"PlayerInput\": 2}"
	// curl -X POST "http://127.0.0.1:8080/start-game-with-computer" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"PlayerInput\": 3}"

	var PlayerInput models.PlayerInputOptions

	requestBody, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(requestBody, &PlayerInput)
	utils.HandleError(err)

	playerInput := PlayerInput.PlayerInput
	playerPick := core.PlayerSelection(playerInput)

	computerPick := core.ComputerSelection()

	winner := core.Game(playerPick, computerPick)

	response := models.ResponseOutput{PlayerPick: playerPick, ComputerPick: computerPick, Winner: winner}
	json.NewEncoder(writer).Encode(response)
}

func startGameWithPlayer(writer http.ResponseWriter, request *http.Request) {
	// curl -X POST "http://127.0.0.1:8080/start-game-with-player" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"PlayerOne\": 1, \"PlayerTwo\": 1}"
	// curl -X POST "http://127.0.0.1:8080/start-game-with-player" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"PlayerOne\": 1, \"PlayerTwo\": 2}"

	var PlayersInput models.TwoPlayerGame
	requestBody, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(requestBody, &PlayersInput)
	utils.HandleError(err)

	playerOneInput := PlayersInput.PlayerOne
	playerTwoInput := PlayersInput.PlayerTwo

	playerOnePick := core.PlayerSelection(playerOneInput)
	playerTwoPick := core.PlayerSelection(playerTwoInput)

	winner := core.Game(playerOnePick, playerTwoPick)

	response := models.ResponseOutputTwoPlayerGame{PlayerOnePick: playerOnePick, PlayerTwoPick: playerTwoPick, Winner: winner}
	json.NewEncoder(writer).Encode(response)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", utils.DisplayGameRules)
	myRouter.HandleFunc("/select-player-or-computer", selectPlayerOrComputer).Methods("POST")
	myRouter.HandleFunc("/start-game-with-computer", startGameWithComputer).Methods("POST")
	myRouter.HandleFunc("/start-game-with-player", startGameWithPlayer).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	utils.PrintText("Serving App: Rock Paper Scissor")
	utils.PrintText("Rest API: Mux Routers")
	utils.PrintText("Running on: http://localhost:8080/")
	handleRequests()
}
