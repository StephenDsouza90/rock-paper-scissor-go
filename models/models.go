package models

// ResponseMessage : A struct for handling responses sent back to the user.
type ResponseMessage struct {
	Message string
}

// PlayerInputStruct : A struct for handling the user input for the JSON request.
type PlayerInputStruct struct {
	UserInput int
}

// TwoPlayerGameStruct : A struct for handling both user input for the JSON request.
type TwoPlayerGameStruct struct {
	UserInputOne int
	UserInputTwo int
}

// GameOutputComputer : A struct for handling the JSON request for the game output between the player and computer.
type GameOutputComputer struct {
	PlayerPick   string
	ComputerPick string
	Winner       string
}

// GameOutputTwoPlayer : A struct for handling the JSON request for the game output between two players.
type GameOutputTwoPlayer struct {
	PlayerOnePick string
	PlayerTwoPick string
	Winner        string
}
