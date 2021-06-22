package helpers

import "math"

func CalculateTournamentPoints(numPlayers, placing int) int {
	if numPlayers < 5 {
		if placing == 1 {
			return 6
		} else if placing == 2 {
			return 3
		} else if placing == 3 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 4 && numPlayers < 7 {
		if placing == 1 {
			return 9
		} else if placing == 2 {
			return 6
		} else if placing == 3 {
			return 3
		} else if placing == 4 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 6 && numPlayers < 9 {
		if placing == 1 {
			return 12
		} else if placing == 2 {
			return 9
		} else if placing == 3 {
			return 6
		} else if placing == 4 {
			return 3
		} else if placing == 5 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 8 && numPlayers < 13 {
		if placing == 1 {
			return 15
		} else if placing == 2 {
			return 12
		} else if placing == 3 {
			return 9
		} else if placing == 4 {
			return 6
		} else if placing == 5 {
			return 3
		} else if placing == 7 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 12 && numPlayers < 17 {
		if placing == 1 {
			return 20
		} else if placing == 2 {
			return 15
		} else if placing == 3 {
			return 12
		} else if placing == 4 {
			return 9
		} else if placing == 5 {
			return 6
		} else if placing == 7 {
			return 3
		} else if placing == 9 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 16 && numPlayers < 25 {
		if placing == 1 {
			return 30
		} else if placing == 2 {
			return 20
		} else if placing == 3 {
			return 15
		} else if placing == 4 {
			return 12
		} else if placing == 5 {
			return 9
		} else if placing == 7 {
			return 6
		} else if placing == 9 {
			return 3
		} else if placing == 13 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 24 && numPlayers < 33 {
		if placing == 1 {
			return 40
		} else if placing == 2 {
			return 30
		} else if placing == 3 {
			return 20
		} else if placing == 4 {
			return 15
		} else if placing == 5 {
			return 12
		} else if placing == 7 {
			return 9
		} else if placing == 9 {
			return 6
		} else if placing == 13 {
			return 3
		} else if placing == 17 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 32 && numPlayers < 49 {
		if placing == 1 {
			return 60
		} else if placing == 2 {
			return 40
		} else if placing == 3 {
			return 30
		} else if placing == 4 {
			return 20
		} else if placing == 5 {
			return 15
		} else if placing == 7 {
			return 12
		} else if placing == 9 {
			return 9
		} else if placing == 13 {
			return 6
		} else if placing == 17 {
			return 3
		} else if placing == 25 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 48 && numPlayers < 65 {
		if placing == 1 {
			return 80
		} else if placing == 2 {
			return 60
		} else if placing == 3 {
			return 40
		} else if placing == 4 {
			return 30
		} else if placing == 5 {
			return 20
		} else if placing == 7 {
			return 15
		} else if placing == 9 {
			return 12
		} else if placing == 13 {
			return 9
		} else if placing == 17 {
			return 6
		} else if placing == 25 {
			return 3
		} else if placing == 33 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 64 && numPlayers < 97 {
		if placing == 1 {
			return 90
		} else if placing == 2 {
			return 70
		} else if placing == 3 {
			return 60
		} else if placing == 4 {
			return 40
		} else if placing == 5 {
			return 30
		} else if placing == 7 {
			return 20
		} else if placing == 9 {
			return 15
		} else if placing == 13 {
			return 12
		} else if placing == 17 {
			return 9
		} else if placing == 25 {
			return 6
		} else if placing == 33 {
			return 3
		} else if placing == 49 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 96 && numPlayers < 129 {
		if placing == 1 {
			return 120
		} else if placing == 2 {
			return 90
		} else if placing == 3 {
			return 70
		} else if placing == 4 {
			return 60
		} else if placing == 5 {
			return 40
		} else if placing == 7 {
			return 30
		} else if placing == 9 {
			return 20
		} else if placing == 13 {
			return 15
		} else if placing == 17 {
			return 12
		} else if placing == 25 {
			return 9
		} else if placing == 33 {
			return 6
		} else if placing == 49 {
			return 3
		} else if placing == 65 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 128 && numPlayers < 193 {
		if placing == 1 {
			return 160
		} else if placing == 2 {
			return 120
		} else if placing == 3 {
			return 90
		} else if placing == 4 {
			return 70
		} else if placing == 5 {
			return 60
		} else if placing == 7 {
			return 40
		} else if placing == 9 {
			return 30
		} else if placing == 13 {
			return 20
		} else if placing == 17 {
			return 15
		} else if placing == 25 {
			return 12
		} else if placing == 33 {
			return 9
		} else if placing == 49 {
			return 6
		} else if placing == 65 {
			return 3
		} else if placing == 97 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 192 && numPlayers < 257 {
		if placing == 1 {
			return 200
		} else if placing == 2 {
			return 160
		} else if placing == 3 {
			return 120
		} else if placing == 4 {
			return 90
		} else if placing == 5 {
			return 70
		} else if placing == 7 {
			return 60
		} else if placing == 9 {
			return 40
		} else if placing == 13 {
			return 30
		} else if placing == 17 {
			return 20
		} else if placing == 25 {
			return 15
		} else if placing == 33 {
			return 12
		} else if placing == 49 {
			return 9
		} else if placing == 65 {
			return 6
		} else if placing == 97 {
			return 3
		} else if placing == 129 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 256 && numPlayers < 385 {
		if placing == 1 {
			return 250
		} else if placing == 2 {
			return 200
		} else if placing == 3 {
			return 160
		} else if placing == 4 {
			return 120
		} else if placing == 5 {
			return 90
		} else if placing == 7 {
			return 70
		} else if placing == 9 {
			return 60
		} else if placing == 13 {
			return 40
		} else if placing == 17 {
			return 30
		} else if placing == 25 {
			return 20
		} else if placing == 33 {
			return 15
		} else if placing == 49 {
			return 12
		} else if placing == 65 {
			return 9
		} else if placing == 97 {
			return 6
		} else if placing == 129 {
			return 3
		} else if placing == 193 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 384 && numPlayers < 513 {
		if placing == 1 {
			return 300
		} else if placing == 2 {
			return 250
		} else if placing == 3 {
			return 200
		} else if placing == 4 {
			return 160
		} else if placing == 5 {
			return 120
		} else if placing == 7 {
			return 90
		} else if placing == 9 {
			return 70
		} else if placing == 13 {
			return 60
		} else if placing == 17 {
			return 40
		} else if placing == 25 {
			return 30
		} else if placing == 33 {
			return 20
		} else if placing == 49 {
			return 15
		} else if placing == 65 {
			return 12
		} else if placing == 97 {
			return 9
		} else if placing == 129 {
			return 6
		} else if placing == 193 {
			return 3
		} else if placing == 257 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 512 && numPlayers < 769 {
		if placing == 1 {
			return 400
		} else if placing == 2 {
			return 300
		} else if placing == 3 {
			return 250
		} else if placing == 4 {
			return 200
		} else if placing == 5 {
			return 160
		} else if placing == 7 {
			return 120
		} else if placing == 9 {
			return 90
		} else if placing == 13 {
			return 70
		} else if placing == 17 {
			return 60
		} else if placing == 25 {
			return 40
		} else if placing == 33 {
			return 30
		} else if placing == 49 {
			return 20
		} else if placing == 65 {
			return 15
		} else if placing == 97 {
			return 12
		} else if placing == 129 {
			return 9
		} else if placing == 193 {
			return 6
		} else if placing == 257 {
			return 3
		} else if placing == 385 {
			return 2
		} else {
			return 1
		}
	} else if numPlayers > 768 && numPlayers < 1025 {
		if placing == 1 {
			return 500
		} else if placing == 2 {
			return 400
		} else if placing == 3 {
			return 300
		} else if placing == 4 {
			return 250
		} else if placing == 5 {
			return 200
		} else if placing == 7 {
			return 160
		} else if placing == 9 {
			return 120
		} else if placing == 13 {
			return 90
		} else if placing == 17 {
			return 70
		} else if placing == 25 {
			return 60
		} else if placing == 33 {
			return 40
		} else if placing == 49 {
			return 30
		} else if placing == 65 {
			return 20
		} else if placing == 97 {
			return 15
		} else if placing == 129 {
			return 12
		} else if placing == 193 {
			return 9
		} else if placing == 257 {
			return 6
		} else if placing == 385 {
			return 3
		} else if placing == 513 {
			return 2
		} else {
			return 1
		}
	} else {
		if placing == 1 {
			return 700
		} else if placing == 2 {
			return 500
		} else if placing == 3 {
			return 400
		} else if placing == 4 {
			return 300
		} else if placing == 5 {
			return 250
		} else if placing == 7 {
			return 200
		} else if placing == 9 {
			return 160
		} else if placing == 13 {
			return 120
		} else if placing == 17 {
			return 90
		} else if placing == 25 {
			return 70
		} else if placing == 33 {
			return 60
		} else if placing == 49 {
			return 40
		} else if placing == 65 {
			return 30
		} else if placing == 97 {
			return 20
		} else if placing == 129 {
			return 15
		} else if placing == 193 {
			return 12
		} else if placing == 257 {
			return 9
		} else if placing == 385 {
			return 6
		} else if placing == 513 {
			return 3
		} else if placing == 769 {
			return 2
		} else {
			return 1
		}
	}
}

func CalculateElo(winnerElo, loserElo int) (updatedWinnerElo, updatedLoserElo int) {
	winnerTransElo := math.Pow10(winnerElo / 400)
	loserTransElo := math.Pow10(loserElo / 400)
	winnerExpectedScore := winnerTransElo / (winnerTransElo + loserTransElo)
	loserExpectedScore := loserTransElo / (winnerTransElo + loserTransElo)
	updatedWinnerElo = int(math.Round(float64(winnerElo) + 32*(1-winnerExpectedScore)))
	updatedLoserElo = int(math.Round(float64(winnerElo) + 32*(0-loserExpectedScore)))
	return
}
