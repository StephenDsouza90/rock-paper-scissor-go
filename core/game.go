package core

var GameOptions map[int]string

func Game(firstPlayerPick, secondPlayerOrComputerPick string) string {
	rock := GameOptions[1]
	paper := GameOptions[2]
	scissor := GameOptions[3]

	var winner string

	if secondPlayerOrComputerPick == firstPlayerPick {
		winner = "Tie"
	} else if secondPlayerOrComputerPick == rock {
		if firstPlayerPick == scissor {
			winner = "Computer/PlayerTwo"
		} else if firstPlayerPick == paper {
			winner = "PlayerOne"
		}
	} else if secondPlayerOrComputerPick == scissor {
		if firstPlayerPick == rock {
			winner = "PlayerOne"
		} else if firstPlayerPick == paper {
			winner = "Computer/PlayerTwo"
		}
	} else if secondPlayerOrComputerPick == paper {
		if firstPlayerPick == rock {
			winner = "Computer/PlayerTwo"
		} else if firstPlayerPick == scissor {
			winner = "PlayerOne"
		}
	} else if secondPlayerOrComputerPick == "" {
		winner = "Computer/PlayerTwo did not select anything"
	} else if firstPlayerPick == "" {
		winner = "PlayerOne did not select anything"
	}

	return winner
}
