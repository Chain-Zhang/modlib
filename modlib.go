package modlib

import(
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func Message() string {
	return "this message from modlib"
}

func MaxIntegerDigits(){
	const year = 1999
	p := message.NewPrinter(language.English)
	p.Println("Year:", number.Decimal(year, number.MaxIntegerDigits(2)))
}

