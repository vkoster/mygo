package main

import (
	"fmt"
	"nested/quote"
	"nested/quote/quotes"
)

func main() {
	fmt.Println("Here are all the quotes I know:")
	fmt.Println(quotes.MyQuotes)
	fmt.Println("...and this is for Volker: ")
	fmt.Println(quote.GetAQuote())
}
