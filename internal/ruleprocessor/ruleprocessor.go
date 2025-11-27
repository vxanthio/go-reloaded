package ruleprocessor

import (
	"strconv"
	"strings"
	"unicode"
)

// ProcessTokens receives the tokens produced by the tokenizer and applies:
// - (hex) → convert previous token from hex to decimal
// - (bin) → convert previous token from binary to decimal
// - (up), (low), (cap) → transform previous token
// - (up,n), (low,n), (cap,n) → transform previous n word tokens
// - a → an before words starting with a vowel or 'h'
// Tags are removed from the output.
// A new slice of tokens is returned.
func ProcessTokens(tokens []string) []string {
	result := make([]string, 0, len(tokens))

	for _, tok := range tokens {
		// If it's not a rule tag, keep it normally
		if !isTag(tok) {
			result = append(result, tok)
			continue
		}

		// Parse the tag: name + optional number
		name, n := parseTag(tok)

		switch name {
		case "hex":
			applyNumberTag(&result, 16, isValidHex)
		case "bin":
			applyNumberTag(&result, 2, isValidBin)
		case "up", "low", "cap":
			applyCaseTag(&result, name, n)
		default:
			// unknown tag → ignored
		}
	}

	// Final pass for the "a → an" rule
	result = applyArticleRule(result)

	return result
}

/////////////////////////////
// Tag Handling
/////////////////////////////

// isTag returns true if token looks like "(something)"
func isTag(tok string) bool {
	if len(tok) < 3 {
		return false
	}
	return tok[0] == '(' && tok[len(tok)-1] == ')'
}

// parseTag converts "(up, 2)" → name="up", n=2
// If no number exists, n defaults to 1.
func parseTag(tok string) (name string, n int) {
	inner := tok[1 : len(tok)-1]
	parts := strings.Split(inner, ",")

	name = strings.TrimSpace(parts[0])
	n = 1 // default

	if len(parts) > 1 {
		if v, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil && v > 0 {
			n = v
		}
	}

	return
}

/////////////////////////////
// Number conversion rules: (hex) / (bin)
/////////////////////////////

// applyNumberTag converts the previous token using the given base (16 or 2).
// validate ensures that the previous token is a valid number for that base.
func applyNumberTag(result *[]string, base int, validate func(string) bool) {
	if len(*result) == 0 {
		return
	}

	prev := (*result)[len(*result)-1]

	// If invalid number, we silently ignore the tag
	if !validate(prev) {
		return
	}

	value, err := strconv.ParseInt(prev, base, 64)
	if err != nil {
		return
	}

	// Replace previous token with decimal string
	(*result)[len(*result)-1] = strconv.FormatInt(value, 10)
}

// Hex validator: 0–9, a–f, A–F
func isValidHex(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !(r >= '0' && r <= '9' ||
			r >= 'a' && r <= 'f' ||
			r >= 'A' && r <= 'F') {
			return false
		}
	}
	return true
}

// Binary validator: only '0' or '1'
func isValidBin(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r != '0' && r != '1' {
			return false
		}
	}
	return true
}

/////////////////////////////
// Case transformation rules
/////////////////////////////

// applyCaseTag applies up/low/cap to the previous n *word* tokens.
// Punctuation tokens are skipped.
func applyCaseTag(result *[]string, mode string, count int) {
	if len(*result) == 0 || count <= 0 {
		return
	}

	remaining := count

	// Walk backward through the token list
	for i := len(*result) - 1; i >= 0 && remaining > 0; i-- {
		tok := (*result)[i]

		if !isWordToken(tok) {
			continue
		}

		switch mode {
		case "up":
			tok = strings.ToUpper(tok)
		case "low":
			tok = strings.ToLower(tok)
		case "cap":
			tok = capitalize(tok)
		}

		(*result)[i] = tok
		remaining--
	}
}

// isWordToken returns true if token contains a letter or a digit.
// Used to skip punctuation tokens.
func isWordToken(tok string) bool {
	for _, r := range tok {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

// capitalize converts word to lowercase and makes first letter uppercase
func capitalize(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

/////////////////////////////
// Article rule: "a" → "an"
/////////////////////////////

// applyArticleRule changes "a" to "an" if the next *word* token
// starts with a vowel or 'h'. Punctuation between them is skipped.
func applyArticleRule(tokens []string) []string {
	if len(tokens) == 0 {
		return tokens
	}

	result := make([]string, len(tokens))
	copy(result, tokens)

	for i := 0; i < len(result); i++ {
		tok := result[i]

		// case-insensitive equality check
		if !strings.EqualFold(tok, "a") {
			continue
		}

		// Find the next word token (skip punctuation tokens)
		j := i + 1
		for j < len(result) && !isWordToken(result[j]) {
			j++
		}
		if j >= len(result) {
			continue
		}

		next := result[j]

		if startsWithVowelOrH(next) {
			if tok == "A" {
				result[i] = "An"
			} else {
				result[i] = "an"
			}
		}
	}

	return result
}

// startsWithVowelOrH returns true if the first letter in the string
// is a vowel (a,e,i,o,u) or 'h'. Leading punctuation/digits are skipped.
func startsWithVowelOrH(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h'
	}
	return false
}
