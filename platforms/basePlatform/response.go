package basePlatform

type Response interface {
	AddText(string) Response
	SetImageCard(string, string, string) Response
	EndSession() Response
	GetBytes() []byte
}
