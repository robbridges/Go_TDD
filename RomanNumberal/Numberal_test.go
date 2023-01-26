package RomanNumberal

import "testing"

func TestRomanNumbers(t *testing.T) {
	cases := []struct {
		description string
		num         int
		want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets convered to III", 3, "III"},
		{"4 gets converted to IV", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
		{"100 gets converted to C", 100, "C"},
		{"90 gets converted to XC", 90, "XC"},
	}

	for _, test := range cases {
		t.Run(test.description, func(t *testing.T) {
			got := ConvertToRoman(test.num)
			want := test.want
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}

		})
	}

}
