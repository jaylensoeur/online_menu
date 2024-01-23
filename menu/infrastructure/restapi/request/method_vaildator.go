package request

import "strings"

func MethodMatched(methods []string, method string) bool {
	var found = false
	for i := range methods {
		if strings.ToUpper(methods[i]) == strings.ToUpper(method) {
			found = true
		}
	}

	return found
}
