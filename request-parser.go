package gospeak

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/blforce/gospeak/platforms/alexa"
	"github.com/blforce/gospeak/platforms/basePlatform"
	"github.com/blforce/gospeak/platforms/dialogflow"
	"github.com/blforce/gospeak/platforms/googleAssistant"
)

func ParseRequestStream(r io.ReadCloser) basePlatform.Request {
	data, _ := ioutil.ReadAll(r)
	defer r.Close()
	defer ioutil.WriteFile("request.json", data, 0644)

	return ParseRequest(data)
}

// ParseRequest turns raw JSON bytes into a Request for whatever assistant platform you're working with
func ParseRequest(data []byte) basePlatform.Request {

	var formatTest map[string]string
	json.Unmarshal(data, &formatTest)

	if _, ok := formatTest["queryResult"]; ok {
		var dialog dialogflow.Request
		json.Unmarshal(data, &dialog)
		return dialog
	}

	_, hasUser := formatTest["user"]
	_, hasConversation := formatTest["conversation"]
	if hasUser && hasConversation {
		var google googleAssistant.Request
		json.Unmarshal(data, &google)
		return google
	}

	var alexa alexa.Request
	json.Unmarshal(data, &alexa)

	return alexa
}
