package utils

import (
	"fmt"
	"os"
	"regexp"
)

// Confirm send the prompt and get result
func Confirm(prompt string) bool {
	var (
		inputStr string
		err      error
	)
	_, err = fmt.Fprint(os.Stdout, prompt)
	if err != nil {
		fmt.Println("fmt.Fprint err", err)
		os.Exit(-1)
	}

	_, err = fmt.Scanf("%s", &inputStr)
	if err != nil {
		os.Exit(-1)
	}

	return getConfirmResult(inputStr)
}

var yesRx = regexp.MustCompile("^(?i:y(?:es)?)$")

// like y|yes|Y|YES return true
func getConfirmResult(str string) bool {
	return yesRx.MatchString(str)
}