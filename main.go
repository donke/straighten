package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/atotto/clipboard"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	re := regexp.MustCompile(`\r?\n`)
	line := re.ReplaceAllString(text, "")

	if err := clipboard.WriteAll(line); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Print(line)
	os.Exit(0)
}
