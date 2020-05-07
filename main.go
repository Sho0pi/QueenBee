package main

import (
	"flag"
	"os"
)

const (
	EmptyString = ""
)

var password string

func init() {
	flag.StringVar(&password, "p", EmptyString, "The Password.")
}

func parametersCheck() {

	if password == EmptyString {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	parametersCheck()

	rule := Rule{}

	//rule.AddRule(emptyRule)
	//rule.AddRule(vowelsToNumbersRule)
	//rule.AddRule(camelMixRule)

	rule.AddRule(momboComboRule)

	rule.Crack([]byte(password))
}
