import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")
	length := len(searchWord)
	for i, w := range words {
		if len(w) < length {
			continue
		}
		if w[:length] == searchWord {
			return i + 1
		}
	}
	return -1
}