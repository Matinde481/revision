package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	n := len(s)
	lastIsUpper := false

	// Validate camelCase rules
	for i := 0; i < n; i++ {
		ch := s[i]

		// Rule: Numbers or punctuation are not allowed anywhere
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) {
			return s
		}

		isUpper := (ch >= 'A' && ch <= 'Z')

		// Rule: No two capitalized letters shall follow directly each other
		if isUpper && lastIsUpper {
			return s
		}
		lastIsUpper = isUpper
	}

	// Rule: The word does not end on a capitalized letter
	if s[n-1] >= 'A' && s[n-1] <= 'Z' {
		return s
	}

	// Build the snake_case version by inserting an underscore before each uppercase letter
	var result []byte
	for i := 0; i < n; i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, ch)
	}

	return string(result)
}
