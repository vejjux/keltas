package schedule

import "strings"

func trim(s string) string {
	return strings.Trim(s, " \n\t*\"")
}
