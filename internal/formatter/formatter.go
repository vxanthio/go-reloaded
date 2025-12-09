package formatter

func Format(tokens []string) string {
	result := ""

	for i, tok := range tokens {

		// 1) Punctuation (.,!?;: or groups like ... !! ?!)
		if isPunctuation(tok) {
			result += tok // attach directly (NO space before)
			continue
		}

		// 2) If it's NOT first token, decide whether to add space
		if i > 0 {
			prev := tokens[i-1]

			// Do NOT put a space if the previous token is a quote '
			if prev != "'" {
				result += " "
			}
		}

		// 3) Add the actual token (word, number, or ')
		result += tok
	}

	return result
}

func isPunctuation(tok string) bool {
	for _, ch := range tok {
		switch ch {
		case '.', ',', '!', '?', ';', ':':
		default:
			return false
		}
	}
	return true
}
