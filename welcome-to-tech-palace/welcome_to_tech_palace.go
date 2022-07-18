package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	countStars := strings.Repeat("*", numStarsPerLine)
	fancyPrint := fmt.Sprintf("%s\n%s\n%s", countStars, welcomeMsg, countStars)
	return fancyPrint

}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanupMessage := strings.ReplaceAll(oldMsg, "*", " ")
	cleanupMessage = strings.TrimSpace(cleanupMessage)
	return cleanupMessage
}
