package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/obynonwane/go-translate-udemy/cli"
)

var sourceLang string
var targetLang string
var sourceText string

func int() {
	flag.StringVar(&sourceLang, "s", "en", "Source language[en]")
	flag.StringVar(&targetLang, "t", "fr", "Target language[fr]")
	flag.StringVar(&sourceText, "st", "", "Text to translate")
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Println("Options")
		os.Exit(1)
	}

	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	cli.RequestTranslate(reqBody)
}
