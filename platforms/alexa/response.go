package alexa

import (
	"encoding/json"
	"fmt"

	"github.com/blforce/gospeak/platforms/basePlatform"
)

type outputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

type image struct {
	SmallImage string `json:"smallImageUrl,omitempty"`
	LargeImage string `json:"largeImageUrl,omitempty"`
}

type card struct {
	Type  string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
	Image *image `json:"image,omitempty"`
}

type responseBody struct {
	OutputSpeech     outputSpeech           `json:"outputSpeech"`
	Card             *card                  `json:"card,omitempty"`
	Reprompt         *outputSpeech          `json:"reprompt,omitempty"`
	Directives       map[string]interface{} `json:"directives,omitempty"`
	ShouldEndSession bool                   `json:"shouldEndSession"`
}

type Response struct {
	Version           string            `json:"version"`
	SessionAttributes map[string]string `json:"sessionAttributes,omitempty"`
	Response          responseBody      `json:"response"`
}

func (r Response) AddText(value string) basePlatform.Response {
	r.Response.OutputSpeech.Type = "PlainText"

	r.Response.OutputSpeech.Text = r.Response.OutputSpeech.Text + value + " "
	return r
}

func (r Response) SetImageCard(title, imageURL, text string) basePlatform.Response {
	r.Response.Card = &card{
		Type:  "Standard",
		Title: title,
		Text:  text,
	}

	if len(imageURL) > 0 {
		r.Response.Card.Image = &image{
			LargeImage: imageURL,
		}
	}

	return r
}

func (r Response) GetBytes() []byte {
	result, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return result
}

func (r Response) EndSession() basePlatform.Response {
	r.Response.ShouldEndSession = true
	return r
}
