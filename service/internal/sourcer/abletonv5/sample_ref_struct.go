package abletonv5

type SampleReferenceDocument struct {
	HasBase

	SampleAbsPath          string `json:"sampleAbsPath,omitempty"`
	SampleFilename         string `json:"sampleFilename,omitempty"`
	SampleOriginalFileSize int64  `json:"sampleOriginalFileSize,omitempty"`
}

func NewSampleReferenceDocument() *SampleReferenceDocument {
	return &SampleReferenceDocument{
		HasBase:                NewHasBase(AbletonSampleReference),
		SampleAbsPath:          "",
		SampleFilename:         "",
		SampleOriginalFileSize: 0,
	}
}
