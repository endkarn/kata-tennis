package tennis_test

import (
	"testing"
	"../tennis"
)

func Test_GetScore_Input_Court_Player_1_Score_0_Court_Player_2_Score_0_Should_Be_LOVE_LOVE(t *testing.T){
	expected := "LOVE-LOVE"
	court := tennis.Court{
		Player1Score:0,
		Player2Score:0,
	}

	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ",expected,actual)
	}
}