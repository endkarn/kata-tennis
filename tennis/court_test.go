package tennis_test

import (
	"testing"
	"../tennis"
)

func Test_GetScore_Input_Court_PlayerOneScore_0_Court_PlayerTwoScore_0_Should_Be_LOVE_LOVE(t *testing.T){
	expected := "LOVE-LOVE"
	court := tennis.Court{
		PlayerOneScore:0,
		PlayerTwoScore:0,
	}

	actual := court.GetScore()

	if expected != actual {
		t.Errorf("Expected in %s but get %s ",expected,actual)
	}
}

