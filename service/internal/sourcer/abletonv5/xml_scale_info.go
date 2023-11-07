package abletonv5

type XmlScaleInformationNode struct {
	ScaleInformation XmlScaleInformationValue `xml:"ScaleInformation"`
}
type XmlScaleInformationValue struct {
	RootNote XmlIntValue    `xml:"RootNote"`
	Name     XmlStringValue `xml:"Name"`
}

func (s *XmlScaleInformationValue) HumanizeRootNote() string {
	switch s.RootNote.Value {
	case 0:
		return "c"
	}

	return "unknown"
}
