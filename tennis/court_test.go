package tennis_test

import "testing"

func Test_GetScore_Input_Court_Player_1_Score_0_Court_Player_2_Score_0_Should_Be_LOVE_LOVE(t *testing.T){
	expected := "LOVE-LOVE"
	court := Court{
		player1Score:0,
		player2Score:0,
	}

	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ",expected,actual)
	}
}