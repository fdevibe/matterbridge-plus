package bridge

import (
	"strings"
)

func tableformatter(nicks_s string, nicksPerRow int) string {
	nicks := strings.Split(nicks_s, " ")
	result := "|IRC users"
	if nicksPerRow < 1 {
		nicksPerRow = 4
	}
	for i := 0; i < 2; i++ {
		for j := 1; j <= nicksPerRow && j <= len(nicks); j++ {
			if i == 0 {
				result += "|"
			} else {
				result += ":-|"
			}
		}
		result += "\r\n|"
	}
	result += nicks[0] + "|"
	for i := 1; i < len(nicks); i++ {
		if i%nicksPerRow == 0 {
			result += "\r\n|" + nicks[i] + "|"
		} else {
			result += nicks[i] + "|"
		}
	}
	return result
}

func plainformatter(nicks string, nicksPerRow int) string {
	return nicks + " currently on IRC"
}

func IsMarkup(message string) bool {
	switch message[0] {
	case '|':
		fallthrough
	case '#':
		fallthrough
	case '_':
		fallthrough
	case '*':
		fallthrough
	case '~':
		fallthrough
	case '-':
		fallthrough
	case ':':
		fallthrough
	case '>':
		fallthrough
	case '=':
		return true
	}
	return false
}
