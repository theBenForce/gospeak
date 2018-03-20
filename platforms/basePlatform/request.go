package basePlatform

const (
	// AmazonAlexa platform
	AmazonAlexa = iota

	// GoogleAssistant platform
	GoogleAssistant = iota

	// Dialogflow platform
	DialogFlow = iota
)

// Request represents messages coming from the assistant platform
type Request interface {
	GetArgument(string) string
	GetArgumentInt(string) int64
	GetIntent() string
	GetRequestType() string
	GetPlatform() int
	GetResponse() Response
	GetLanguage() string
}
