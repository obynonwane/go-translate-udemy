// Package cli
package cli

import "net/http"

// Custom struct for body from cli input
type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

// google translate url
const translateUrl = "https://translate.googleapis.com/translate_a/single"

// RequestTranslate for making request to google api to translate a text to a gicen languate
func RequestTranslate(body *RequestBody) {

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

	//make request to translate url
	client.Do(req)
}
