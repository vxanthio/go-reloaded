package tokenizer

func Tokenize(text string) []string {
	currentPunctuation := ""
	currentword := ""
	tokens := []string{}
	currentRule := ""
	isInRule := false
	for _, ch := range text {
		if ch == '(' {
			isInRule = true
		}
		if isInRule == true {
			if ch != '(' && ch != ')' {
				currentRule += string(ch)
			}
		}
		if isInRule == false {
			if ch == ')' {
				isInRule = false
				tokens = append(tokens, currentRule)
				currentRule = ""
				if ch != '(' {
					currentRule += string(ch)
				}
				continue
			}
		}
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
