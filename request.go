package gospeak

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/blforce/gospeakAlexa"
	"github.com/blforce/gospeakCommon"
	"github.com/blforce/gospeakDialogflow"
	"github.com/blforce/gospeakGoogleAssistant"
)

func ParseRequestStream(r io.ReadCloser) gospeakCommon.Request {
	data, _ := ioutil.ReadAll(r)

	defer r.Close()

	return ParseRequest(data)
}

// ParseRequest turns raw JSON bytes into a Request for whatever assistant platform you're working with
func ParseRequest(data []byte) gospeakCommon.Request {

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
