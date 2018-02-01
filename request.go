package gospeak

import "encoding/json"

const (
	// AmazonAlexa platform
	AmazonAlexa = iota

	// GoogleAssistant platform
	GoogleAssistant = iota
)

type Request interface {
	GetSlot(string) string
	GetIntent() string
	GetRequestType() string
}

func ParseRequest(data []byte) Request {
	var request alexaRequest
	json.Unmarshal(data, &request)

	return request
}
