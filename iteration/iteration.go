package iteration

import (
	"errors"
	"strings"
)

func IterateCharacter(str string, count int) (string, error) {
	if count < 0 {
		return "Count must be a positive value", errors.New("Invalid count value")
	}

	var iteratedString = ""

	for i := 1; i <= count; i++ {
		iteratedString += str
	}

	return iteratedString, nil
}

// ReplaceCharacter A replace function to replace characters within a string, you pass the string you want to replace, the match you want
// replaced and what you want it replaced with the value -1 is passed in to indicate we want to replace all matches instead
// of n
func ReplaceCharacter(origionalString, replaced, replacedBy string) string {
	replacedString := strings.Replace(origionalString, replaced, replacedBy, -1)
	return replacedString
}
