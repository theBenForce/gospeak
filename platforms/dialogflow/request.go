package dialogflow

import (
	"fmt"
	"strconv"

	"github.com/blforce/gospeak/platforms/basePlatform"
)

type context struct {
	Name          string            `json:"name"`
	LifespanCount uint16            `json:"lifespanCount"`
	Parameters    map[string]string `json:"parameters"`
}

type part struct {
	Text        string `json:"text"`
	EntityType  string `json:"entityType"`
	Alias       string `json:"alias"`
	UserDefined bool   `json:"userDefined"`
}

type trainingPhrase struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Parts      []part `json:"parts"`
	TimesAdded uint32 `json:"timesAddedCount"`
}

type parameter struct {
	Name                  string   `json:"name"`
	DisplayName           string   `json:"displayName"`
	Value                 string   `json:"value"`
	DefaultValue          string   `json:"defaultValue"`
	EntityTypeDisplayName string   `json:"entityTypeDisplayName"`
	Mandatory             bool     `json:"mandatory"`
	Prompts               []string `json:"prompts"`
	IsList                bool     `json:"isList"`
}

type followupIntentInfo struct {
}

type intent struct {
	Name              string           `json:"name"`
	DisplayName       string           `json:"displayName"`
	Priority          uint16           `json:"priority"`
	IsFallback        bool             `json:"isFallbak"`
	MlEnabled         bool             `json:"mlEnabled"`
	InputContextNames []string         `json:"inputContextNames"`
	Events            []string         `json:"events"`
	TrainingPhrases   []trainingPhrase `json:"trainingPhrases"`
	Action            string           `json:"action"`
	OutputContexts    []context        `json:"outputContexts"`
	ResetContexts     bool             `json:"resetContexts"`
}

type queryResult struct {
	QueryText                   string                 `json:"queryText"`
	LanguageCode                string                 `json:"languageCode"`
	SpeechRecognitionConfidence float32                `json:"speechRecognitionConfidence"`
	Action                      string                 `json:"action"`
	Parameters                  map[string]interface{} `json:"parameters"`
	AllRequiredPramsPresent     bool                   `json:"allRequiredParamsPresent"`
	FulfillmentText             string                 `json:"fulfillmentText"`
	FulfillmentMessages         []interface{}          `json:"fulfillmentMessages"`
	WebhookSource               string                 `json:"webhookSource"`
	WebhookPayload              interface{}            `json:"webhookPayload"`
	OutputContexts              []context              `json:"outputContexts"`
	Intent                      intent                 `json:"intent"`
	IntentDetectionConfidence   float32                `json:"intentDetectionConfidence"`
	DiagnosticInfo              interface{}            `json:"diagnostricInfo"`
}

type originalDetectIntentRequest struct {
	Source  string      `json:"source"`
	Payload interface{} `json:"payload"`
}

type Request struct {
	Session                     string                      `json:"session"`
	ResponseID                  string                      `json:"responseId"`
	QueryResult                 queryResult                 `json:"queryResult"`
	OriginalDetectIntentRequest originalDetectIntentRequest `json:"originalDetectIntentRequest"`
}

func (r Request) GetRequestType() string {
	return "IntentRequest"
}

func (r Request) GetArgument(slot string) string {
	if val, ok := r.QueryResult.Parameters[slot]; ok {
		output := fmt.Sprintf("%v", val)
		return output
	}

	return ""
}

func (r Request) GetArgumentInt(slot string) int64 {
	result, _ := strconv.ParseInt(r.GetArgument(slot), 10, 64)
	return result
}

func (r Request) GetIntent() string {
	return r.QueryResult.Action
}

func (r Request) GetPlatform() int {
	return basePlatform.DialogFlow
}

func (r Request) GetResponse() basePlatform.Response {
	return Response{
		request: r,
	}
}

func (r Request) GetLanguage() string {
	return r.QueryResult.LanguageCode
}
