package tennis

type Court struct {
	PlayerOneScore int
	PlayerTwoScore int
}

func (court Court) GetScore() string {
	if court.PlayerTwoScore == 1 {
		return "LOVE-"+"15"
	}
	if court.PlayerTwoScore == 2 {
		return "LOVE-"+"30"
	}
	if court.PlayerTwoScore == 3 {
		return "LOVE-"+"40"
	}
	if court.PlayerTwoScore == 4 {
		return "Player 2 Win"
	}
	return "LOVE-LOVE"
}
