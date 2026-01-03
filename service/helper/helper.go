package helper

import (
	"regexp"
	"strings"
	"unicode"
)

func ContainsSpecialStrict(s string) bool {
	for _, char := range s {
		// If it is NOT a letter AND NOT a number, it is considered special
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return true
		}
	}
	return false
}

// IsPossibleSQLInjection checks for common SQL injection patterns.
// WARNING: This is a heuristic (guess).
// 1. It may flag innocent text (False Positive).
// 2. It is NOT a replacement for Parameterized Queries ($1, $2).
func IsPossibleSQLInjection(s string) bool {
	// Normalize to uppercase for case-insensitive checking
	upperS := strings.ToUpper(s)

	// List of dangerous patterns often used in attacks
	patterns := []string{
		`--`,                 // SQL line comment
		`/\*.*\*/`,           // SQL block comment
		`;`,                  // Query chaining (Stacking)
		`UNION\s+SELECT`,     // Union based injection
		`OR\s+\d+=\d+`,       // Tautology (e.g., OR 1=1)
		`AND\s+\d+=\d+`,      // Tautology
		`DROP\s+TABLE`,       // Destructive command
		`DELETE\s+FROM`,      // Destructive command
		`INSERT\s+INTO`,      // Destructive command
		`UPDATE\s+\w+\s+SET`, // Destructive command
		`XP_CMDSHELL`,        // MSSQL specific execution
	}

	for _, pattern := range patterns {
		// regexp.MatchString allows us to find the pattern anywhere in the string
		matched, _ := regexp.MatchString(pattern, upperS)
		if matched {
			return true
		}
	}

	return false
}
