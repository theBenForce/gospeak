package gospeak

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blforce/gospeakCommon"
)

type IntentHandler func(gospeakCommon.Request) gospeakCommon.Response

type Handler struct {
	handlers map[string]IntentHandler
}

func NewHandler() Handler {
	result := Handler{}
	result.handlers = make(map[string]IntentHandler)

	return result
}

func (h Handler) RegisterIntentHandler(intentName string, handler IntentHandler) {
	h.handlers[intentName] = handler
}

func (h Handler) ExecuteRequest(req gospeakCommon.Request) gospeakCommon.Response {
	method := h.handlers[req.GetIntent()]

	return method(req)
}

func (h Handler) HandleWebRequest(w http.ResponseWriter, r *http.Request) {
	req := ParseRequestStream(r.Body)

	response := h.ExecuteRequest(req)
	w.WriteHeader(http.StatusOK)
	w.Write(response.GetBytes())
}

func (h Handler) HandleLambdaRequest(ctx context.Context, event map[string]interface{}) (string, error) {
	var eventBody string

	if body, ok := event["body"]; ok {
		eventBody = fmt.Sprintf("%s", body)
	} else {
		original, err := json.Marshal(event)

		if err != nil {
			panic(err)
		}

		eventBody = string(original)
	}

	req := ParseRequest([]byte(eventBody))

	response := h.ExecuteRequest(req)
	return string(response.GetBytes()), nil
}
