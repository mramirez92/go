package blackjack

import "testing" 

func TestStrToInt(t *testing.T){

	got := ParseCard("seven")
	want := 7

	if got != want{
		t.Errorf("got %q, wanted %q", got, want)
	}
}