package main

func evaluate(s string, knowledge [][]string) string {
	result := ""
	curStart := 0
	started := false
	kMap := make(map[string]string)
	for _, k := range knowledge {
		kMap[k[0]] = k[1]
	}

	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			var replaced string
			if r, ok := kMap[s[curStart+1:i]]; ok {
				replaced = r
			} else {
				replaced = "?"
			}
			result += replaced
			started = false
			continue
		}

		if started {
			continue
		}

		if s[i] == '(' {
			started = true
			curStart = i
			continue
		}
		result += string(s[i])
	}
	return result
}
