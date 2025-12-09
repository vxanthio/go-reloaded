package formatter

// Format takes the processed tokens and rebuilds the final text.
// It adds spaces only where needed and keeps punctuation and quotes correctly placed.
func Format(tokens []string) string {
	result := ""

	for i, tok := range tokens {

		// 1. Standard punctuation: . , ! ? ; :
		if isPunctuation(tok) {
			// NEVER put a space before punctuation
			result += tok
			continue
		}

		// 2. Single quote: '
		if tok == "'" {
			// If next token is a word, then this quote is OPENING
			isOpening := i+1 < len(tokens) &&
				!isPunctuation(tokens[i+1]) &&
				tokens[i+1] != "'"

			if isOpening {
				// For an opening quote, add one space BEFORE it (if needed)
				// Example: Wilson: 'I am...
				if len(result) > 0 && result[len(result)-1] != ' ' {
					result += " "
				}
			}
			// Closing quote: attach directly (no extra spaces)
			result += "'"
			continue
		}

		// 3. Normal word or number
		if i > 0 {
			prev := tokens[i-1]

			// If the previous token was an opening quote, DO NOT add a space.
			// We want 'I (no space between quote and word)
			if prev != "'" {
				result += " "
			}
		}

		result += tok
	}

	return result
}

// isPunctuation checks if a token contains ONLY punctuation from the set
// . , ! ? ; : (quotes are NOT included here!)
func isPunctuation(tok string) bool {
	for _, ch := range tok {
		switch ch {
		case '.', ',', '!', '?', ';', ':':
		// allowed punctuation, keep checking
		default:
			return false // found a non-punctuation character
		}
	}
	return true // every character was valid punctuation
}
