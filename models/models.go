package models

type PlayWithOptions struct {
	PlayWith int
}

type PlayerInputOptions struct {
	PlayerInput int
}

type TwoPlayerGame struct {
	PlayerOne int
	PlayerTwo int
}

type ResponseOutput struct {
	PlayerPick   string
	ComputerPick string
	Winner       string
}

type ResponseOutputTwoPlayerGame struct {
	PlayerOnePick string
	PlayerTwoPick string
	Winner        string
}

type Message struct {
	Message string
}
