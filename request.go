package gospeak

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/blforce/gospeakAlexa"
	"github.com/blforce/gospeakDialogflow"
	"github.com/blforce/gospeakGoogleAssistant"
)

const (
	// AmazonAlexa platform
	AmazonAlexa = 0

	// GoogleAssistant platform
	GoogleAssistant = 1

	// Dialogflow platform
	DialogFlow = 2
)

// Request represents messages coming from the assistant platform
type Request interface {
	GetArgument(string) string
	GetIntent() string
	GetRequestType() string
	GetPlatform() int
	GetResponse() Response
}

type Response interface {
	SetText(string) Response
	GetBytes() []byte
}

func ParseRequestStream(r io.ReadCloser) Request {
	data := ioutil.ReadAll(r)
	return ParseRequest(data)
}

// ParseRequest turns raw JSON bytes into a Request for whatever assistant platform you're working with
func ParseRequest(data []byte) Request {

	var formatTest map[string]string
	json.Unmarshal(data, &formatTest)

	if _, ok := formatTest["queryResult"]; ok {
		var dialog gospeakDialogflow.Request
		json.Unmarshal(data, &dialog)
		return dialog
	}

	_, hasUser := formatTest["user"]
	_, hasConversation := formatTest["conversation"]
	if hasUser && hasConversation {
		var google gospeakGoogleAssistant.Request
		json.Unmarshal(data, &google)
		return google
	}

	var alexa gospeakAlexa.Request
	json.Unmarshal(data, &alexa)

	return alexa
}
