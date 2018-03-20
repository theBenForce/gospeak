package alexa

import (
	"strconv"

	"github.com/blforce/gospeak/platforms/basePlatform"
)

type application struct {
	ID string `json:"applicationId"`
}

type permissions struct {
	ConsentToken string `json:"consentToken"`
}

type user struct {
	ID          string      `json:"userId"`
	AccessToken string      `json:"accessToken"`
	Permissions permissions `json:"permissions"`
}

type session struct {
	New         bool                   `json:"new"`
	SessionID   string                 `json:"sessionId"`
	Application application            `json:"application"`
	Attributes  map[string]interface{} `json:"attributes"`
	User        user                   `json:"user"`
}

type device struct {
	ID                  string                 `json:"deviceId"`
	SupportedInterfaces map[string]interface{} `json:"supportedInterfaces"`
}

type system struct {
	Device         device      `json:"device"`
	Application    application `json:"application"`
	User           user        `json:"user"`
	APIEndpoint    string      `json:"apiEndpoint"`
	APIAccessToken string      `json:"apiAccessToken"`
}

type audioPlayer struct {
	PlayerActivity string `json:"playerActivity"`
	Token          string `json:"token"`
	Offset         uint64 `json:"offsetInMilliseconds"`
}

type context struct {
	System      system      `json:"System"`
	AudioPlayer audioPlayer `json:"AudioPlayer"`
}

type alexaError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type slotResolution struct {
	Authority string `json:"authority"`
	Status    struct {
		Code string `json:"code"`
	} `json:"status"`
	Values []struct {
		Value struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"value"`
	} `json:"values"`
}

type slot struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	ConfirmationStatus string `json:"confirmationStatus"`
	Resolutions        *struct {
		ResolutionsPerAuthority []slotResolution `json:"resolutionsPerAuthority,omitempty"`
	} `json:"resolutions,omitempty"`
}

type intent struct {
	Name               string          `json:"name"`
	ConfirmationStatus string          `json:"confirmationStatus"`
	Slots              map[string]slot `json:"slots"`
}

type requestBody struct {
	Type                 string     `json:"type"`
	ID                   string     `json:"requestId"`
	Timestamp            string     `json:"timestamp"`
	Token                string     `json:"token"`
	OffsetInMilliseconds uint64     `json:"offsetInMilliseconds"`
	Locale               string     `json:"locale"`
	Reason               string     `json:"reason"`
	Error                alexaError `json:"error"`
	DialogState          string     `json:"dialogState"`
	Intent               intent     `json:"intent"`
}

// AlexaRequest structure that will be sent from Alexa
type Request struct {
	Version string      `json:"version"`
	Session session     `json:"session"`
	Context context     `json:"context"`
	Request requestBody `json:"request"`
}

func (r Request) GetRequestType() string {
	return r.Request.Type
}

func getBestValue(val slot) string {
	result := val.Value

	if val.Resolutions != nil {
		if len(val.Resolutions.ResolutionsPerAuthority) > 0 {
			resolution := val.Resolutions.ResolutionsPerAuthority[0]

			if len(resolution.Values) > 0 && resolution.Status.Code == "ER_SUCCESS_MATCH" {
				result = resolution.Values[0].Value.Name
			}
		}
	}

	return result
}

func (r Request) GetArgument(slot string) string {
	if r.Request.Type != "IntentRequest" {
		return ""
	}

	if val, ok := r.Request.Intent.Slots[slot]; ok {
		return getBestValue(val)
	}

	return ""
}

func (r Request) GetArgumentInt(slot string) int64 {
	result, _ := strconv.ParseInt(r.GetArgument(slot), 10, 64)
	return result
}

func (r Request) GetIntent() string {
	if r.Request.Type != "IntentRequest" {
		return "_" + r.Request.Type
	}

	return r.Request.Intent.Name
}

func (r Request) GetPlatform() int {
	return basePlatform.AmazonAlexa
}

func (r Request) GetResponse() basePlatform.Response {
	return Response{
		Version: "1.0",
	}
}

func (r Request) GetLanguage() string {
	return r.Request.Locale
}
