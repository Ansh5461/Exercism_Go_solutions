package techpalace
import (
	"strings"
    )
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	s := "Welcome to the Tech Palace, "+strings.ToUpper(customer)
    return s
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var i int
    var s string
    for i=0;i<numStarsPerLine;i++ {
        s = s+ "*"
    }
	s = s+ "\n"+welcomeMsg+"\n"
    for i=0;i<numStarsPerLine;i++ {
        s = s+ "*"
    }
	return s
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    var s string
    s = strings.Trim(oldMsg, "*")
    s = strings.TrimSpace(s)
    s = strings.Trim(s, "*")
    s = strings.TrimSpace(s)
	return s
}
