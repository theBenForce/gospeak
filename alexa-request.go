package gospeak

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
	DeviceID            string                 `json:"deviceId"`
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

type alexaSlotResolution struct {
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

type alexaSlot struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	ConfirmationStatus string `json:"confirmationStatus"`
	Resolutions        struct {
		ResolutionsPerAuthority []alexaSlotResolution `json:"resolutionsPerAuthority"`
	} `json:"resolutions"`
}

type alexaIntent struct {
	Name               string               `json:"name"`
	ConfirmationStatus string               `json:"confirmationStatus"`
	Slots              map[string]alexaSlot `json:"slots"`
}

type requestBody struct {
	Type                 string      `json:"type"`
	ID                   string      `json:"requestId"`
	Timestamp            string      `json:"timestamp"`
	Token                string      `json:"token"`
	OffsetInMilliseconds uint64      `json:"offsetInMilliseconds"`
	Locale               string      `json:"locale"`
	Reason               string      `json:"reason"`
	Error                alexaError  `json:"error"`
	DialogState          string      `json:"dialogState"`
	Intent               alexaIntent `json:"intent"`
}

// Request structure that will be sent from Alexa
type alexaRequest struct {
	Version string      `json:"version"`
	Session session     `json:"session"`
	Context context     `json:"context"`
	Request requestBody `json:"request"`
}

func (r alexaRequest) GetRequestType() string {
	return r.Request.Type
}

func (r alexaRequest) GetSlot(slot string) string {
	if r.Request.Type != "IntentRequest" {
		return ""
	}

	if val, ok := r.Request.Intent.Slots[slot]; ok {
		return val.Value
	}

	return ""
}

func (r alexaRequest) GetIntent() string {
	if r.Request.Type != "IntentRequest" {
		return r.Request.Type
	}

	return r.Request.Intent.Name
}
