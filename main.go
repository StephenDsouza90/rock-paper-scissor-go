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

func init() {
	// By default, the init() is called before the main().
	// The game options are global so that a player or computer has the same source for selecting options.

	core.GameOptions = make(map[int]string)
	core.GameOptions[0] = "Rock"
	core.GameOptions[1] = "Paper"
	core.GameOptions[2] = "Scissor"
}

func selectPlayerOrComputer(writer http.ResponseWriter, request *http.Request) {
	// In this func, users can select to play with either another player or with a computer.

	// 1 is for player
	// 2 is for computer
	const playWithPlayer = 1
	const playWithComputer = 2

	// PlayWith variable holds the input from the user.
	var PlayWith models.PlayerInputStruct
	requestBody, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(requestBody, &PlayWith)
	utils.HandleError(err)

	// Action based on user input.
	userInput := PlayWith.UserInput
	if userInput == playWithPlayer {
		responseOne := models.ResponseMessage{Message: "You are now playing with another player"}
		responseTwo := models.ResponseMessage{Message: "Select one of the following: Press 1 for Rock, Press 2 for Paper, Press 3 for Scissor, Press 0 to exit"}
		json.NewEncoder(writer).Encode(responseOne)
		json.NewEncoder(writer).Encode(responseTwo)
	} else if userInput == playWithComputer {
		responseOne := models.ResponseMessage{Message: "You are now playing with the computer"}
		responseTwo := models.ResponseMessage{Message: "Select one of the following: Press 1 for Rock, Press 2 for Paper, Press 3 for Scissor, Press 0 to exit"}
		json.NewEncoder(writer).Encode(responseOne)
		json.NewEncoder(writer).Encode(responseTwo)
	}
}

func startGameWithComputer(writer http.ResponseWriter, request *http.Request) {
	// In this func, a player will play against a computer.
	// The computer will select an option at random.

	// PlayerInput variable holds the input from the user.
	var PlayerInput models.PlayerInputStruct
	requestBody, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(requestBody, &PlayerInput)
	utils.HandleError(err)

	// Choice of the player
	playerInput := PlayerInput.UserInput
	playerPick := core.GetPlayerSelection(playerInput)

	// Choice of the computer
	computerPick := core.GetComputerSelection()

	winner := core.PlayGame(playerPick, computerPick)

	response := models.GameOutputComputer{PlayerPick: playerPick, ComputerPick: computerPick, Winner: winner}
	json.NewEncoder(writer).Encode(response)
}

func startGameWithPlayer(writer http.ResponseWriter, request *http.Request) {
	// In this func, two users can play with each other.
	// Both players provide the input in the json body.

	// PlayersInput variable holds both the input from the users.
	var PlayersInput models.TwoPlayerGameStruct
	requestBody, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(requestBody, &PlayersInput)
	utils.HandleError(err)

	// Get player's one pick
	playerOneInput := PlayersInput.UserInputOne
	playerOnePick := core.GetPlayerSelection(playerOneInput)

	// Get player's two pick
	playerTwoInput := PlayersInput.UserInputTwo
	playerTwoPick := core.GetPlayerSelection(playerTwoInput)

	winner := core.PlayGame(playerOnePick, playerTwoPick)

	response := models.GameOutputTwoPlayer{PlayerOnePick: playerOnePick, PlayerTwoPick: playerTwoPick, Winner: winner}
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
	utils.PrintText("Running on: http://127.0.0.1:8080/")
	handleRequests()
}
