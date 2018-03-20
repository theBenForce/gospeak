package googleAssistant

import (
	"strconv"

	"github.com/blforce/gospeak/platforms/basePlatform"
)

type userProfile struct {
	DisplayName string `json:"displayName"`
	GivenName   string `json:"givenName"`
	FamilyName  string `json:"familyName"`
}

type user struct {
	ID          string      `json:"userId"`
	Profile     userProfile `json:"profile"`
	AccessToken string      `json:"accessToken"`
	Permissions []string    `json:"permissions"`
	Locale      string      `json:"locale"`
	LastSeen    string      `json:"userStorage"`
}

type postalAddress struct {
	Revision           int      `json:"revision"`
	RegionCode         string   `json:"regionCode"`
	LanguageCode       string   `json:"languageCode"`
	PostalCode         string   `json:"postalCode"`
	SortingCode        string   `json:"sortingCode"`
	AdministrativeArea string   `json:"administrativeArea"`
	Locality           string   `json:"locality"`
	SubLocality        string   `json:"subLocality"`
	AddressLines       []string `json:"addressLines"`
	Recipients         []string `json:"recipients"`
	Organization       string   `json:"organization"`
}

type location struct {
	Coordinates struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"coordinates"`
	FormattedAddress string        `json:"formattedAddress"`
	ZipCode          string        `json:"zipCode"`
	City             string        `json:"city"`
	PostalAddress    postalAddress `json:"postalAddress"`
	Name             string        `json:"name"`
	PhoneNumber      string        `json:"phoneNumber"`
	Notes            string        `json:"notes"`
}

type deviceCapability struct {
	Name string `json:"name"`
}

type conversation struct {
	ID    string `json:"conversationId"`
	Type  string `json:"type"`
	Token string `json:"conversationToken"`
}

type rawInput struct {
	InputType string `json:"inputType"`
	Query     string `json:"query"`
}

type dateTime struct {
	Date struct {
		Year  uint16 `json:"year"`
		Month uint8  `json:"month"`
		Day   uint8  `json:"day"`
	} `json:"date"`
	Time struct {
		Hours   uint8 `json:"hours"`
		Minutes uint8 `json:"minutes"`
		Seconds uint8 `json:"seconds"`
		Nanos   uint  `json:"nanos"`
	} `json:"time"`
}

type argument struct {
	Name          string      `json:"name"`
	RawText       string      `json:"rawText"`
	TextValue     string      `json:"textValue"`
	BoolValue     bool        `json:"boolValue"`
	DateTimeValue dateTime    `json:"datetimeValue"`
	Extension     interface{} `json:"extension"`
}

type input struct {
	RawInputs []rawInput `json:"rawInputs"`
	Intent    string     `json:"intent"`
	Arguments []argument `json:"arguments"`
}

// Request structure that will be sent from Google Assistant
type Request struct {
	User   user `json:"user"`
	Device struct {
		Location location `json:"location"`
	} `json:"device"`
	Surface struct {
		Capabilities []deviceCapability `json:"capabilities"`
	} `json:"surface"`
	Conversation conversation `json:"conversation"`
	Inputs       []input      `json:"inputs"`
	IsInSandbox  bool         `json:"isInSandbox"`
}

func (r Request) GetRequestType() string {
	return "IntentRequest"
}

func (r Request) GetArgument(slot string) string {
	return ""
}

func (r Request) GetArgumentInt(slot string) int64 {
	result, _ := strconv.ParseInt(r.GetArgument(slot), 10, 64)
	return result
}

func (r Request) GetIntent() string {
	return ""
}

func (r Request) GetPlatform() int {
	return basePlatform.GoogleAssistant
}

func (r Request) GetResponse() basePlatform.Response {
	return nil
}

func (r Request) GetLanguage() string {
	return ""
}
