package tennis

const winablescore = 4

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
	if court.PlayerTwoScore == winablescore {
		return "Player 2 Win"
	}
	if court.PlayerOneScore == winablescore {
		return "Player 1 Win"
	}
	return scoreMap[court.PlayerOneScore] + "-" + scoreMap[court.PlayerTwoScore]
}
