package tennis_test

import (
	"../tennis"
	"testing"
)

func Test_GetScore_Input_Court_PlayerOneScore_0_Court_PlayerTwoScore_0_Should_Be_LOVE_LOVE(t *testing.T) {
	expected := "LOVE-LOVE"
	court := tennis.NewCourt()

	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ", expected, actual)
	}
}

func Test_GetScore_Input_Court_PlayerOneScore_0_Court_PlayerTwoScore_1_Should_Be_LOVE_15(t *testing.T) {
	expected := "LOVE-15"
	court := tennis.NewCourt()

	court.WinRoundPlayer(2)
	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ", expected, actual)
	}
}

func Test_GetScore_Input_Court_PlayerOneScore_0_Court_PlayerTwoScore_2_Should_Be_LOVE_30(t *testing.T) {
	expected := "LOVE-30"
	court := tennis.NewCourt()

	court.WinRoundPlayer(2)
	court.WinRoundPlayer(2)
	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ", expected, actual)
	}
}

func Test_GetScore_Input_Court_PlayerOneScore_0_Court_PlayerTwoScore_3_Should_Be_LOVE_40(t *testing.T) {
	expected := "LOVE-40"
	court := tennis.NewCourt()

	court.WinRoundPlayer(2)
	court.WinRoundPlayer(2)
	court.WinRoundPlayer(2)
	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ", expected, actual)
	}
}

func Test_GetScore_Input_Court_PlayerOneScore_0_Court_PlayerTwoScore_4_Should_Be_Player_2_Win(t *testing.T) {
	expected := "Player 2 Win"
	court := tennis.NewCourt()

	court.WinRoundPlayer(2)
	court.WinRoundPlayer(2)
	court.WinRoundPlayer(2)
	court.WinRoundPlayer(2)
	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ", expected, actual)
	}
}

func Test_GetScore_Input_Court_PlayerOneScore_1_Court_PlayerTwoScore_0_Should_Be_15_LOVE(t *testing.T) {
	expected := "15-LOVE"
	court := tennis.NewCourt()

	court.WinRoundPlayer(1)
	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ", expected, actual)
	}
}

func Test_GetScore_Input_Court_PlayerOneScore_2_Court_PlayerTwoScore_0_Should_Be_15_LOVE(t *testing.T) {
	expected := "30-LOVE"
	court := tennis.NewCourt()

	court.WinRoundPlayer(1)
	court.WinRoundPlayer(1)
	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ", expected, actual)
	}
}

func Test_GetScore_Input_Court_PlayerOneScore_3_Court_PlayerTwoScore_0_Should_Be_15_LOVE(t *testing.T) {
	expected := "40-LOVE"
	court := tennis.NewCourt()

	court.WinRoundPlayer(1)
	court.WinRoundPlayer(1)
	court.WinRoundPlayer(1)
	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ", expected, actual)
	}
}

func Test_GetScore_Input_Court_PlayerOneScore_4_Court_PlayerTwoScore_0_Should_Be_15_LOVE(t *testing.T) {
	expected := "Player 1 Win"
	court := tennis.NewCourt()

	court.WinRoundPlayer(1)
	court.WinRoundPlayer(1)
	court.WinRoundPlayer(1)
	court.WinRoundPlayer(1)
	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ", expected, actual)
	}
}

func Test_GetScore_Input_Court_PlayerOneScore_2_Court_PlayerTwoScore_2_Should_Be_30_30(t *testing.T) {
	expected := "30-30"
	court := tennis.NewCourt()

	court.WinRoundPlayer(1)
	court.WinRoundPlayer(2)
	court.WinRoundPlayer(1)
	court.WinRoundPlayer(2)
	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ", expected, actual)
	}
}
