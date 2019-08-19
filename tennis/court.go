package tennis

type Court struct {
	PlayerOneScore int
	PlayerTwoScore int
}

func (court Court) GetScore() string {
	scoreMap := map[int]string{
		0: "LOVE",
		1: "15",
		2: "30",
		3: "40",
	}
	if court.PlayerTwoScore == 4 {
		return "Player 2 Win"
	}
	if court.PlayerOneScore == 4 {
		return "Player 1 Win"
	}
	return scoreMap[court.PlayerOneScore] + "-" + scoreMap[court.PlayerTwoScore]
}
