package main

import (
	"testing"

	"github.com/StephenDsouza90/RockPaperScissor/core"
)

func TestGetPlayerSelection(t *testing.T) {
	// Rock
	playerPickRock := core.GetPlayerSelection(0)
	if playerPickRock != "Rock" {
		t.Error("Player picked 0 which is supposed to be Rock")
	}

	// Paper
	playerPickPaper := core.GetPlayerSelection(1)
	if playerPickPaper != "Paper" {
		t.Error("Player picked 1 which is supposed to be Paper")
	}

	// Scissor
	playerPickScissor := core.GetPlayerSelection(2)
	if playerPickScissor != "Scissor" {
		t.Error("Player picked 2 which is supposed to be Scissor")
	}
}

func TestPlayGame(t *testing.T) {
	var tests = []struct {
		playerOneInput string
		playerTwoInput string
		expectedWinner string
	}{
		{"Rock", "Paper", "Computer/PlayerTwo"},    // Paper covers rock
		{"Rock", "Scissor", "PlayerOne"},           // Rock smashes scissor
		{"Scissor", "Rock", "Computer/PlayerTwo"},  // Rock smashes scissor
		{"Scissor", "Paper", "PlayerOne"},          // Scissor cuts paper
		{"Paper", "Rock", "PlayerOne"},             // Paper covers rock
		{"Paper", "Scissor", "Computer/PlayerTwo"}, // Scissor cuts paper
		{"Paper", "Paper", "Tie"},                  // Tie
	}

	for _, test := range tests {
		winner := core.PlayGame(test.playerOneInput, test.playerTwoInput) // (test.input)
		if winner != test.expectedWinner {
			t.Error("Test Failed: {} playerOneInput, {} playerTwoInput, expectedWinner: {} actualWinner {}",
				test.playerOneInput, test.playerTwoInput, test.expectedWinner, winner)
		}
	}

}
