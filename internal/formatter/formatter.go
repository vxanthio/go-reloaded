package formatter

// Format takes the processed tokens and rebuilds the final text.
// It handles punctuation, spacing, and special quote placement.
func Format(tokens []string) string {
	result := ""
	insideQuotes := false // Track whether we are between two ' '

	for i, tok := range tokens {

		// ---------------------------------
		// 1) Standard punctuation . , ! ? ; :
		// ---------------------------------
		if isPunctuation(tok) {
			result += tok // ALWAYS attach directly
			continue
		}

		// ---------------------------------
		// 2) Single quote '
		// ---------------------------------
		if tok == "'" {

			// If we are not inside quotes → it is an OPENING quote
			if !insideQuotes {

				// Add a space before opening quote if needed
				if len(result) > 0 && result[len(result)-1] != ' ' {
					result += " "
				}

				result += "'"       // place opening quote
				insideQuotes = true // now we are inside
				continue
			}

			// Otherwise → CLOSING quote
			result += "'"        // attach directly (no extra space)
			insideQuotes = false // we closed it
			continue
		}

		// ---------------------------------
		// 3) Regular words or numbers
		// ---------------------------------
		if i > 0 {
			prev := tokens[i-1]

			// DO NOT insert space after opening quote
			if prev != "'" {
				result += " "
			}
		}

		result += tok
	}

	return result
}

// isPunctuation checks if token contains ONLY punctuation from the set
// . , ! ? ; : (quotes are NOT included here!)
func isPunctuation(tok string) bool {
	for _, ch := range tok {
		switch ch {
		case '.', ',', '!', '?', ';', ':':
		// allowed punctuation
		default:
			return false
		}
	}
	return true
}
