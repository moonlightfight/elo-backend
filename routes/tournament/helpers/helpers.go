package helpers

import "math"

func CalculateTournamentPoints(numPlayers, placing int) int {
	score := math.Abs(float64(placing - numPlayers))
	return int(score)
}
