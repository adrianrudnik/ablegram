package abletonsrc

type InfoTextDocument struct {
	HasBase
	HasUserInfoText
	Parent string `json:"parent,omitempty"`
}

func NewInfoTextDocument() *InfoTextDocument {
	return &InfoTextDocument{
		HasBase:         NewHasBase(AbletonInfoText),
		HasUserInfoText: NewHasUserInfoText(),
	}
}
