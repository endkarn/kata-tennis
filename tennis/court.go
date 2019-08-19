package tennis

type Court struct {
	Player1Score int
	Player2Score int
}

func (court Court) GetScore() string {
	return "LOVE-LOVE"
}
