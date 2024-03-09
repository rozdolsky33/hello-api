package translation

import "strings"

func Translate(word string, language string) string {
	word = sanitizeInput(word)
	language = sanitizeInput(language)

	if word != "hello" {
		return ""
	}

	switch language {
	case "english":
		return "hello"
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	default:
		return ""
	}
}

func sanitizeInput(word string) string {
	word = strings.ToLower(word) //hello
	return strings.TrimSpace(word)
}
