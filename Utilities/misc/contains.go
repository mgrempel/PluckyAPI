package misc

import (
	"strings"
)

//Contains .contains function for slices. Returns true if a slice contains a string
func Contains(slice []string, search string) bool {
	for _, currentString := range slice {
		if strings.Compare(currentString, search) == 0 {
			return true
		}
	}
	return false
}
