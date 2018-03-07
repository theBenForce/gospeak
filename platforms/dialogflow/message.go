package dialogflow

type button struct {
	Text     string `json:"text,omitempty"`
	Postback string `json:"postback,omitempty"`
}

type simpleResponse struct {
	TextToSpeech string `json:"textToSpeech,omitempty"`
	SSML         string `json:"ssml,omitempty"`
	DisplayText  string `json:"displayText,omitempty"`
}

type simpleResponses struct {
	SimpleResponses []simpleResponse `json:"simpleResponses"`
}

type text struct {
	Text []string `json:"text,omitempty"`
}

type image struct {
	ImageURI string `json:"imageUri,omitempty"`
}

type quickReplies struct {
	Title        string   `json:"title,omitempty"`
	QuickReplies []string `json:"quickReplies,omitempty"`
}

type card struct {
	Title    string   `json:"title,omitempty"`
	Subtitle string   `json:"subtitle,omitempty"`
	ImageURI string   `json:"imageUri,omitempty"`
	Buttons  []button `json:"buttons,omitempty"`
}

type basicCard struct {
	Title         string   `json:"title,omitempty"`
	Subtitle      string   `json:"subtitle,omitempty"`
	FormattedText string   `json:"formattedText,omitempty"`
	Image         *image   `json:"image,omitempty"`
	Buttons       []button `json:"buttons,omitempty"`
}

type suggestions struct {
	Suggestions []struct {
		Title string `json:"suggestion"`
	} `json:"suggestions"`
}

type message struct {
	Platform        string           `json:"platform,omitempty"`
	Text            *text            `json:"text,omitempty"`
	Image           *image           `json:"image,omitempty"`
	QuickReplies    *quickReplies    `json:"quickReplies,omitempty"`
	Card            *card            `json:"card,omitempty"`
	Payload         *interface{}     `json:"payload,omitempty"`
	SimpleResponses *simpleResponses `json:"simpleResponses,omitempty"`
	BasicCard       *basicCard       `json:"basicCard,omitempty"`
	Suggestions     *suggestions     `json:"suggestions,omitempty"`
}
