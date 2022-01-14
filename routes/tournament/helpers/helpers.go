package helpers

import (
	"math"
)

func CalculateTournamentPoints(numPlayers, placing int) int {
	switch {
	case numPlayers < 5:
		switch placing {
		case 1:
			return 6
		case 2:
			return 3
		case 3:
			return 2
		default:
			return 1
		}
	case numPlayers > 4 && numPlayers < 7:
		switch placing {
		case 1:
			return 9
		case 2:
			return 6
		case 3:
			return 3
		case 4:
			return 2
		default:
			return 1
		}
	case numPlayers > 6 && numPlayers < 9:
		switch placing {
		case 1:
			return 12
		case 2:
			return 9
		case 3:
			return 6
		case 4:
			return 3
		case 5:
			return 2
		default:
			return 1
		}
	case numPlayers > 8 && numPlayers < 13:
		switch placing {
		case 1:
			return 15
		case 2:
			return 12
		case 3:
			return 9
		case 4:
			return 6
		case 5:
			return 3
		case 7:
			return 2
		default:
			return 1
		}
	case numPlayers > 12 && numPlayers < 17:
		switch placing {
		case 1:
			return 20
		case 2:
			return 15
		case 3:
			return 12
		case 4:
			return 9
		case 5:
			return 6
		case 7:
			return 3
		case 9:
			return 2
		default:
			return 1
		}
	case numPlayers > 16 && numPlayers < 25:
		switch placing {
		case 1:
			return 30
		case 2:
			return 20
		case 3:
			return 15
		case 4:
			return 12
		case 5:
			return 9
		case 7:
			return 6
		case 9:
			return 3
		case 13:
			return 2
		default:
			return 1
		}
	case numPlayers > 24 && numPlayers < 33:
		switch placing {
		case 1:
			return 40
		case 2:
			return 30
		case 3:
			return 20
		case 4:
			return 15
		case 5:
			return 12
		case 7:
			return 9
		case 9:
			return 6
		case 13:
			return 3
		case 17:
			return 2
		default:
			return 1
		}
	case numPlayers > 32 && numPlayers < 49:
		switch placing {
		case 1:
			return 60
		case 2:
			return 40
		case 3:
			return 30
		case 4:
			return 20
		case 5:
			return 15
		case 7:
			return 12
		case 9:
			return 9
		case 13:
			return 6
		case 17:
			return 3
		case 25:
			return 2
		default:
			return 1
		}
	case numPlayers > 48 && numPlayers < 65:
		switch placing {
		case 1:
			return 80
		case 2:
			return 60
		case 3:
			return 40
		case 4:
			return 30
		case 5:
			return 20
		case 7:
			return 15
		case 9:
			return 12
		case 13:
			return 9
		case 17:
			return 6
		case 25:
			return 3
		case 33:
			return 2
		default:
			return 1
		}
	case numPlayers > 48 && numPlayers < 65:
		switch placing {
		case 1:
			return 80
		case 2:
			return 60
		case 3:
			return 40
		case 4:
			return 30
		case 5:
			return 20
		case 7:
			return 15
		case 9:
			return 12
		case 13:
			return 9
		case 17:
			return 6
		case 25:
			return 3
		case 33:
			return 2
		default:
			return 1
		}
	case numPlayers > 64 && numPlayers < 97:
		switch placing {
		case 1:
			return 90
		case 2:
			return 70
		case 3:
			return 60
		case 4:
			return 40
		case 5:
			return 30
		case 7:
			return 20
		case 9:
			return 15
		case 13:
			return 12
		case 17:
			return 9
		case 25:
			return 6
		case 33:
			return 3
		case 49:
			return 2
		default:
			return 1
		}
	case numPlayers > 96 && numPlayers < 129:
		switch placing {
		case 1:
			return 120
		case 2:
			return 90
		case 3:
			return 70
		case 4:
			return 60
		case 5:
			return 40
		case 7:
			return 30
		case 9:
			return 20
		case 13:
			return 15
		case 17:
			return 12
		case 25:
			return 9
		case 33:
			return 6
		case 49:
			return 3
		case 65:
			return 2
		default:
			return 1
		}
	case numPlayers > 128 && numPlayers < 193:
		switch placing {
		case 1:
			return 160
		case 2:
			return 120
		case 3:
			return 90
		case 4:
			return 70
		case 5:
			return 60
		case 7:
			return 40
		case 9:
			return 30
		case 13:
			return 20
		case 17:
			return 15
		case 25:
			return 12
		case 33:
			return 9
		case 49:
			return 6
		case 65:
			return 3
		case 97:
			return 2
		default:
			return 1
		}
	case numPlayers > 192 && numPlayers < 257:
		switch placing {
		case 1:
			return 200
		case 2:
			return 160
		case 3:
			return 120
		case 4:
			return 90
		case 5:
			return 70
		case 7:
			return 60
		case 9:
			return 40
		case 13:
			return 30
		case 17:
			return 20
		case 25:
			return 15
		case 33:
			return 12
		case 49:
			return 9
		case 65:
			return 6
		case 97:
			return 3
		case 129:
			return 2
		default:
			return 1
		}
	case numPlayers > 256 && numPlayers < 385:
		switch placing {
		case 1:
			return 250
		case 2:
			return 200
		case 3:
			return 160
		case 4:
			return 120
		case 5:
			return 90
		case 7:
			return 70
		case 9:
			return 60
		case 13:
			return 40
		case 17:
			return 30
		case 25:
			return 20
		case 33:
			return 15
		case 49:
			return 12
		case 65:
			return 9
		case 97:
			return 6
		case 129:
			return 3
		case 193:
			return 2
		default:
			return 1
		}
	case numPlayers > 384 && numPlayers < 513:
		switch placing {
		case 1:
			return 300
		case 2:
			return 250
		case 3:
			return 200
		case 4:
			return 160
		case 5:
			return 120
		case 7:
			return 90
		case 9:
			return 70
		case 13:
			return 60
		case 17:
			return 40
		case 25:
			return 30
		case 33:
			return 20
		case 49:
			return 15
		case 65:
			return 12
		case 97:
			return 9
		case 129:
			return 6
		case 193:
			return 3
		case 257:
			return 2
		default:
			return 1
		}
	case numPlayers > 512 && numPlayers < 769:
		switch placing {
		case 1:
			return 400
		case 2:
			return 300
		case 3:
			return 250
		case 4:
			return 200
		case 5:
			return 160
		case 7:
			return 120
		case 9:
			return 90
		case 13:
			return 70
		case 17:
			return 60
		case 25:
			return 40
		case 33:
			return 30
		case 49:
			return 20
		case 65:
			return 15
		case 97:
			return 12
		case 129:
			return 9
		case 193:
			return 6
		case 257:
			return 3
		case 385:
			return 2
		default:
			return 1
		}
	case numPlayers > 768 && numPlayers < 1025:
		switch placing {
		case 1:
			return 500
		case 2:
			return 400
		case 3:
			return 300
		case 4:
			return 250
		case 5:
			return 200
		case 7:
			return 160
		case 9:
			return 120
		case 13:
			return 90
		case 17:
			return 70
		case 25:
			return 60
		case 33:
			return 40
		case 49:
			return 30
		case 65:
			return 20
		case 97:
			return 15
		case 129:
			return 12
		case 193:
			return 9
		case 257:
			return 6
		case 385:
			return 3
		case 513:
			return 2
		default:
			return 1
		}
	default:
		switch placing {
		case 1:
			return 700
		case 2:
			return 500
		case 3:
			return 400
		case 4:
			return 300
		case 5:
			return 250
		case 7:
			return 200
		case 9:
			return 160
		case 13:
			return 120
		case 17:
			return 90
		case 25:
			return 70
		case 33:
			return 60
		case 49:
			return 40
		case 65:
			return 30
		case 97:
			return 20
		case 129:
			return 15
		case 193:
			return 12
		case 257:
			return 9
		case 385:
			return 6
		case 513:
			return 3
		case 769:
			return 2
		default:
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
