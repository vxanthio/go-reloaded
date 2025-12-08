package tokenizer

func Tokenize(text string) []string {
	currentWord := ""
	currentPunctuation := ""
	currentRule := ""
	isInRule := false
	tokens := []string{}

	for _, ch := range text {

		// -------------------------
		// RULE START
		// -------------------------
		if ch == '(' {
			// Flush what existed before entering rule
			if currentWord != "" {
				tokens = append(tokens, currentWord)
				currentWord = ""
			}
			if currentPunctuation != "" {
				tokens = append(tokens, currentPunctuation)
				currentPunctuation = ""
			}

			isInRule = true
			currentRule = ""
			continue
		}

		// -------------------------
		// INSIDE RULE
		// -------------------------
		if isInRule {
			if ch == ')' {
				isInRule = false
				tokens = append(tokens, "("+currentRule+")")
				currentRule = ""
				continue
			}
			currentRule += string(ch)
			continue
		}

		// -------------------------
		// GROUPED PUNCTUATION LOGIC (... !! ?? !? ?!)
		// -------------------------
		if ch == '.' || ch == '!' || ch == '?' {
			if currentPunctuation != "" {
				last := currentPunctuation[len(currentPunctuation)-1]

				// ...
				if ch == '.' && last == '.' {
					currentPunctuation += string(ch)
					continue
				}
				// !!
				if ch == '!' && last == '!' {
					currentPunctuation += string(ch)
					continue
				}
				// !?
				if ch == '!' && last == '?' {
					currentPunctuation += string(ch)
					continue
				}
				// ?!
				if ch == '?' && last == '!' {
					currentPunctuation += string(ch)
					continue
				}
				// ??
				if ch == '?' && last == '?' {
					currentPunctuation += string(ch)
					continue
				}
			}
		}

		// -------------------------
		// WORD DETECTION (letters/digits)
		// -------------------------
		if (ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') {

			// flush punctuation before starting a new word
			if currentPunctuation != "" {
				tokens = append(tokens, currentPunctuation)
				currentPunctuation = ""
			}

			currentWord += string(ch)
			continue
		}

		// -------------------------
		// SINGLE PUNCTUATION
		// -------------------------
		if ch == '.' || ch == ',' || ch == '!' || ch == '?' || ch == ':' || ch == ';' || ch == '\'' {
			if currentWord != "" {
				tokens = append(tokens, currentWord)
				currentWord = ""
			}
			currentPunctuation += string(ch)
			continue
		}

		// -------------------------
		// SPACE
		// -------------------------
		if ch == ' ' {
			if currentWord != "" {
				tokens = append(tokens, currentWord)
				currentWord = ""
			}
			if currentPunctuation != "" {
				tokens = append(tokens, currentPunctuation)
				currentPunctuation = ""
			}
			continue
		}
	}

	// -------------------------
	// FINAL FLUSH AFTER LOOP
	// -------------------------
	if currentWord != "" {
		tokens = append(tokens, currentWord)
	}
	if currentPunctuation != "" {
		tokens = append(tokens, currentPunctuation)
	}

	return tokens
}
