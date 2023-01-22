package iteration

import "testing"

func Test_Iterate_Character(t *testing.T) {
	want := "ababababab"
	got, err := IterateCharacter("ab", 5)
	if err != nil {
		t.Errorf("An error was raised with an appropriate value was passed")
	}

	if want != got {
		t.Errorf("Iteration output does not match, expected %s but got %s", want, got)
	}
}

func Test_Iterate_Character_Sad(t *testing.T) {
	want := "Count must be a positive value"
	got, err := IterateCharacter("ab", -1)
	if err == nil {
		t.Errorf("There whould have been an error on a failed value ")
	}
	if want != got {
		t.Error("Count must be a positive value")
	}
}

func Test_Replace_Character_Happy(t *testing.T) {
	want := "moo moo moo"
	got := ReplaceCharacter("oink oink oink", "oink", "moo")
	if want != got {
		t.Errorf("We did not get the correct output when replacing the strings, it should have been %s, but it was %s", want, got)
	}
}
