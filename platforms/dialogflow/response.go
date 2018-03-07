package dialogflow

import (
	"encoding/json"
	"fmt"

	"github.com/blforce/gospeak/platforms/basePlatform"
)

type eventInput struct {
	Name         string                 `json:"name,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
	LanguageCode string                 `json:"languageCode,omitempty"`
}

type Response struct {
	FulfillmentText     string       `json:"fulfillmentText,omitempty"`
	FulfillmentMessages []message    `json:"fulfillmentMessages,omitempty"`
	Source              string       `json:"source,omitempty"`
	Payload             *interface{} `json:"payload,omitempty"`
	OutputContexts      []context    `json:"outputContexts,omitempty"`
	FollowupEventInput  *interface{} `json:"followupEventInput,omitempty"`
	platform            string
	request             Request
}

func (r Response) SetPlatform(newPlatform string) Response {
	if r.platform == newPlatform {
		return r
	}

	r.platform = newPlatform

	for idx := range r.FulfillmentMessages {
		r.FulfillmentMessages[idx].Platform = newPlatform
	}

	return r
}

func (r Response) AddText(value string) basePlatform.Response {
	simpleMessage := simpleResponse{
		DisplayText:  value,
		TextToSpeech: value,
	}

	message := message{
		Platform:        r.platform,
		SimpleResponses: &simpleResponses{},
	}

	message.SimpleResponses.SimpleResponses = append(message.SimpleResponses.SimpleResponses, simpleMessage)

	r.FulfillmentMessages = append(r.FulfillmentMessages, message)

	r.FulfillmentText = r.FulfillmentText + value + " "

	return r
}

func (r Response) SetImageCard(title, imageURL, text string) basePlatform.Response {
	cardMessage := message{
		Platform: r.platform,
		BasicCard: &basicCard{
			Title:    title,
			Subtitle: text,
		},
	}

	if len(imageURL) > 0 {
		cardMessage.BasicCard.Image = &image{
			ImageURI: imageURL,
		}
	}

	r.FulfillmentMessages = append(r.FulfillmentMessages, cardMessage)

	return r
}

func (r Response) GetBytes() []byte {
	result, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return result
}

func (r Response) EndSession() basePlatform.Response {
	// TODO: Figure out if I can end the session from this response
	return r
}
