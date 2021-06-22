package helpers

import "math"

func CalculateTournamentPoints(numPlayers, placing int) int {
	score := math.Abs(float64(placing - numPlayers))
	if score == 0 {
		score++
	}
	return int(score)
}
