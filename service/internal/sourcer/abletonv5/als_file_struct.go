package abletonv5

type AlsFileDocument struct {
	HasBase
}

func NewAlsFileDocument() *AlsFileDocument {
	return &AlsFileDocument{
		HasBase: NewHasBase(AbletonAlsFile),
	}
}
