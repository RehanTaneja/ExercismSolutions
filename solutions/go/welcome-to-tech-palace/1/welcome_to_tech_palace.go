package techpalace

import ("strings";
       "fmt"
       )

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	name := strings.ToUpper(customer)
    return fmt.Sprintf("Welcome to the Tech Palace, %v",name)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	msg := strings.Repeat("*",numStarsPerLine)+"\n"+welcomeMsg+"\n"+strings.Repeat("*",numStarsPerLine)
    return msg
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msg := strings.ReplaceAll(oldMsg,"*","")
    msg = strings.TrimSpace(msg)
    return msg
}
