package tennis

const winablescore = 4

type Court struct {
	PlayerOneScore int
	PlayerTwoScore int
}

func NewCourt() Court {
	return Court{
		PlayerOneScore: 0,
		PlayerTwoScore: 0,
	}
}

func (court *Court) WinRoundPlayer(player int) {
	if player == 1 {
		court.PlayerOneScore++
	} else {
		court.PlayerTwoScore++
	}
}

func (court Court) GetScore() string {
	scoreMap := map[int]string{
		0: "LOVE",
		1: "15",
		2: "30",
		3: "40",
	}
	winner, haveWinner := court.GetWinner()
	if haveWinner {
		return winner
	}
	return scoreMap[court.PlayerOneScore] + "-" + scoreMap[court.PlayerTwoScore]
}

func (court Court) GetWinner() (string, bool) {
	if court.PlayerTwoScore == winablescore {
		return "Player 2 Win", true
	}
	if court.PlayerOneScore == winablescore {
		return "Player 1 Win", true
	}
	return "No Winner", false
}
