package gospeak

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/blforce/gospeakCommon"
	"github.com/blforce/src/github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-lambda-go/events"
)

type RequestHandler func(gospeakCommon.Request) gospeakCommon.Response

type Handler struct {
	handlers map[string]RequestHandler
	aliases  map[string]string
}

func NewHandler() Handler {
	result := Handler{}
	result.handlers = make(map[string]RequestHandler)
	result.aliases = make(map[string]string)

	return result
}

func (h Handler) RegisterIntentHandler(intentName string, handler RequestHandler) {
	h.handlers[intentName] = handler
}

func (h Handler) Redirect(intent, target string) {
	h.aliases[intent] = target
}

func (h Handler) RedirectLaunch(target string) {
	h.Redirect("_LaunchRequest", target)
}

func (h Handler) RegisterUnhandled(handler RequestHandler) {
	h.RegisterIntentHandler("Unhandled", handler)
}

func (h Handler) RegisterLaunchHandler(handler RequestHandler) {
	h.RegisterIntentHandler("_LaunchRequest", handler)
}

func (h Handler) RegisterSessionEndedHandler(handler RequestHandler) {
	h.RegisterIntentHandler("_SessionEndedRequest", handler)
}

func (h Handler) getRequestIntent(req gospeakCommon.Request) string {
	originalIntent := req.GetIntent()

	if alias, ok := h.aliases[originalIntent]; ok {
		return alias
	}

	return originalIntent
}

func (h Handler) ExecuteRequest(req gospeakCommon.Request) gospeakCommon.Response {
	if method, ok := h.handlers[h.getRequestIntent(req)]; ok {
		return method(req)
	} else if method, ok := h.handlers["Unhandled"]; ok {
		return method(req)
	} else {
		fmt.Printf("Unknown intent \"%s\" and no handler was specified for \"Unhandled\".\n", req.GetIntent())
		return req.GetResponse().AddText(fmt.Sprintf("Unknown intent: %s", req.GetIntent()))
	}
}

func (h Handler) Start() {

	environment := os.Getenv("GOSPEAK_ENVIRONMENT")

	if len(environment) == 0 {
		environment = "LAMBDA"
	}

	if environment == "HTTP" {
		path := os.Getenv("GOSPEAK_HTTP_PATH")
		if len(path) == 0 {
			path = "/"
		}

		port := os.Getenv("GOSPEAK_HTTP_PORT")
		if len(port) == 0 {
			port = "8000"
		}

		http.HandleFunc(path, h.HandleWebRequest)

		fmt.Printf("GoSpeak is listening on port %s at path %s\n", port, path)
		http.ListenAndServe(":"+port, nil)
	} else {
		lambda.Start(h.HandleLambdaRequest)
	}
}

func (h Handler) HandleWebRequest(w http.ResponseWriter, r *http.Request) {
	req := ParseRequestStream(r.Body)

	response := h.ExecuteRequest(req)
	w.Header().Set("Content-Type", "application/json")
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
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: string(responseBody.GetBytes()),
		}
	} else {
		response = responseBody
	}

	return response, nil
}
