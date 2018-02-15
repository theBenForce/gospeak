package gospeak

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blforce/gospeakCommon"

	"github.com/aws/aws-lambda-go/events"
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
	if method, ok := h.handlers[req.GetIntent()]; ok {
		return method(req)
	} else {
		fmt.Printf("Unknown intent: %s\n", req.GetIntent())
		return req.GetResponse().SetText(fmt.Sprintf("Unknown intent: %s", req.GetIntent()))
	}
}

func (h Handler) HandleWebRequest(w http.ResponseWriter, r *http.Request) {
	req := ParseRequestStream(r.Body)

	response := h.ExecuteRequest(req)
	w.WriteHeader(http.StatusOK)
	w.Write(response.GetBytes())
}

func (h Handler) HandleLambdaRequest(ctx context.Context, event map[string]interface{}) (interface{}, error) {
	var eventBody string
	isApiGateway := false

	if body, ok := event["body"]; ok {
		eventBody = fmt.Sprintf("%s", body)
		isApiGateway = true
	} else {
		original, err := json.Marshal(event)

		if err != nil {
			panic(err)
		}

		eventBody = string(original)
	}

	req := ParseRequest([]byte(eventBody))

	var response interface{}
	responseBody := h.ExecuteRequest(req)

	if isApiGateway {

		response = events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(responseBody.GetBytes()),
		}
	} else {
		response = responseBody
	}

	return response, nil
}
