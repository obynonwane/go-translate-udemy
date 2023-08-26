// Package cli
package cli

import (
	"log"
	"net/http"
	"sync"

	"github.com/Jeffail/gabs"
)

// Custom struct for body from cli input
type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

// google translate url
const translateUrl = "https://translate.googleapis.com/translate_a/single"

// RequestTranslate for making request to google api to translate a text to a gicen languate
func RequestTranslate(body *RequestBody, str chan string, wg *sync.WaitGroup) {

	// create http client
	client := &http.Client{}

	//initialize your request
	req, err := http.NewRequest("GET", translateUrl, nil)

	//extract your query params and create your query params for the request
	query := req.URL.Query()
	query.Add("client", "gtx")
	query.Add("sl", body.SourceLang)
	query.Add("tl", body.TargetLang)
	query.Add("dt", "t")
	query.Add("q", body.SourceText)

	//encode your query params back to json
	req.URL.RawQuery = query.Encode()

	if err != nil {
		log.Fatal("1. There was a problem: %s", err)
	}
	//make request to translate url
	res, err := client.Do(req)

	if err != nil {
		log.Fatal("2. There was a problem: %s", err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		str <- "You have been rate limited, try again later"
		wg.Done()
		return
	}

	//parse to json
	parsedJson, err := gabs.ParseJSONBuffer(res.Body)
	if err != nil {
		log.Fatal("3. There was a problem: %s", err)
	}

	//denest the parsed json
	nestOne, err := parsedJson.ArrayElement(0)
	if err != nil {
		log.Fatal("4. There was a problem: %s", err)
	}
	nestTwo, err := nestOne.ArrayElement(0)
	if err != nil {
		log.Fatal("5. There was a problem: %s", err)
	}

	translatedString, err := nestTwo.ArrayElement(0)
	if err != nil {
		log.Fatal("6. There was a problem: %s", err)
	}

	str <- translatedString.Data().(string)

	//end goroutine
	wg.Done()
}
