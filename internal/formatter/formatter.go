package formatter

func Format(tokens []string) string {
	result := ""

	for i, tok := range tokens {

		// If it's punctuation, attach DIRECTLY (no space before it)
		if isPunctuation(tok) {
			result += tok
			continue
		}

		// If it's a word/token (not punctuation) and it's NOT the first token → add space
		if i > 0 {
			result += " "
		}

		// Now add the actual word/token
		result += tok
	}

	return result
}

func isPunctuation(tok string) bool {
	for _, ch := range tok {
		switch ch {
		case '.', ',', '!', '?', ';', ':':
		// It's allowed punctuation → keep checking others
		default:
			return false // token contains something else than punctuation
		}
	}
	return true // All characters were punctuation
}
