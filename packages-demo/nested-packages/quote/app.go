package quote

import (
	"math/rand"
	"nested/quote/quotes"
)

// Zufallszahl für ein Quote generieren
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func GetAQuote() string {
	return quotes.MyQuotes[random(0, len(quotes.MyQuotes))]
}
