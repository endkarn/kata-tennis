package tennis

type Court struct {
	PlayerOneScore int
	PlayerTwoScore int
}



func (court Court) GetScore() string {
	scoreMap := map[int]string{
		0:"LOVE",
		1:"15",
		2:"30",
		3:"40",
	}

	var score string
	score = "LOVE-"
	if court.PlayerTwoScore == 1 {
		return score+scoreMap[1]
	}
	if court.PlayerTwoScore == 2 {
		return score+scoreMap[2]
	}
	if court.PlayerTwoScore == 3 {
		return score+scoreMap[3]
	}
	if court.PlayerTwoScore == 4 {
		return "Player 2 Win"
	}
	return "LOVE-LOVE"
}
