package tennis

type Court struct {
	PlayerOneScore int
	PlayerTwoScore int
}

func (court Court) GetScore() string {
	return "LOVE-LOVE"
}
