package main

import (
	"flag" //parsing cli arguements
	"fmt"  //handling output
	"os"   //handle cli operations
	"strings"
	"sync" //

	"github.com/obynonwane/go-translate-udemy/cli" //custom package
)

// define a wait group to have a list of goroutine processes
// so main function would not end before the processes
var wg = sync.WaitGroup{}

// declaration of global variables - to hold values entered via the cli
var sourceLang string
var targetLang string
var sourceText string

/*
special function in golang that get executed before the main function
this function was used to define cli flags (-s, -t, -st respectively)
this flags allows user to specify source, target language and the text to translate
*/
func init() {
	flag.StringVar(&sourceLang, "s", "en", "Source language[en]")
	flag.StringVar(&targetLang, "t", "fr", "Target language[fr]")
	flag.StringVar(&sourceText, "st", "", "Text to translate")
}

// Main function entry point for our golang cli application
func main() {
	//processes the cli args and assign their value to the global variable
	flag.Parse()

	//counts number of flags that where set if none it prints the options and exits the cli
	if flag.NFlag() == 0 {
		fmt.Println("Options")
		os.Exit(1)
	}

	//define a channel of string - for comunication between goroutine and your main process
	// telling th emain process when the goroutine process is finished
	strChan := make(chan string)

	//add goroutine to a waitgroup
	wg.Add(1)

	//return a pointer to a struct RequestBody in cli package from main package in main function
	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	//make a call to RequestTranslate function in cli package - making it a goroutine
	go cli.RequestTranslate(reqBody, strChan, &wg)

	//publishing goroutine response to channel for main process to acknoledge
	processedStr := strings.ReplaceAll(<-strChan, "+", " ")
	fmt.Printf("%s\n", processedStr)
	close(strChan)

	//wait for the goroutine to finish
	wg.Wait()
}
