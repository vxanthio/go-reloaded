package tokenizer

func Tokenize(text string) []string {
	currentPunctuation := ""
	currentword := ""
	tokens := []string{}
	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' {
			currentword += string(ch)
			if currentPunctuation != "" {
				tokens = append(tokens, currentPunctuation)
				currentPunctuation = ""
			}
		}

		if ch == ':' || ch == ';' || ch == ',' || ch == '.' || ch == '!' || ch == '?' {
			if currentword != "" {
				tokens = append(tokens, currentword)
			}
			currentword = ""
			currentPunctuation += string(ch)
		}
		if ch == ' ' {
			if currentword != "" {
				tokens = append(tokens, currentword)
				currentword = ""
			}
			if currentPunctuation != "" {
				tokens = append(tokens, currentPunctuation)
				currentPunctuation = ""
			}
		}

	}
	if currentword != "" {
		tokens = append(tokens, currentword)
		currentword = ""
	}
	if currentPunctuation != "" {
		tokens = append(tokens, currentPunctuation)
		currentPunctuation = ""
	}
	return tokens

}
