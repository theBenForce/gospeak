package gospeak

import (
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