package core

// GameOptions : Options of the game.
var GameOptions map[int]string

// PlayGame : The main business logic of the game.
func PlayGame(firstPlayerPick, secondPlayerPick string) string {
	// In case of playing with computer then secondPlayerPick is the computer.

	rock := GameOptions[0]
	paper := GameOptions[1]
	scissor := GameOptions[2]

	var winner string

	if secondPlayerPick == firstPlayerPick {
		winner = "Tie"
	} else if secondPlayerPick == rock {
		if firstPlayerPick == scissor {
			winner = "Computer/PlayerTwo"
		} else if firstPlayerPick == paper {
			winner = "PlayerOne"
		}
	} else if secondPlayerPick == scissor {
		if firstPlayerPick == rock {
			winner = "PlayerOne"
		} else if firstPlayerPick == paper {
			winner = "Computer/PlayerTwo"
		}
	} else if secondPlayerPick == paper {
		if firstPlayerPick == rock {
			winner = "Computer/PlayerTwo"
		} else if firstPlayerPick == scissor {
			winner = "PlayerOne"
		}
	} else if secondPlayerPick == "" {
		winner = "Computer/PlayerTwo did not select anything"
	} else if firstPlayerPick == "" {
		winner = "PlayerOne did not select anything"
	}

	return winner
}
