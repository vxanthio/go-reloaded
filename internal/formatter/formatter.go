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
			//If next token exists and is NOT punctuation and is NOT "'",
			//then we add ONE space.
			if i+1 < len(tokens) && !isPunctuation(tokens[i-1]) && tokens[i+1] != "'" {
				result += " "
			}

			continue
		}
		//If "'" , attach with no space
		if tok == "'" {
			result += tok
			continue
		}
		// If previous token was "'",attach with no space
		if i > 0 && tokens[i-1] != "'" {
			result += " "
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
