package gospeak

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlexaIntentRequestParsing(t *testing.T) {
	buf, _ := ioutil.ReadFile("example-requests/alexa/intent-request.json")

	request := ParseRequest(buf)

	assert.EqualValues(t, "IntentRequest", request.GetRequestType())
	assert.EqualValues(t, "GetZodiacHoroscopeIntent", request.GetIntent())

	slotValue := request.GetSlot("ZodiacSign")
	assert.EqualValues(t, "virgo", slotValue)
}
