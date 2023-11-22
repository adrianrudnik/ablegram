package abletonsrc

type XmlTempoNode struct {
	Manual XmlFloatValue `xml:"Manual"`
}

type XmlTempoWithToggleNode struct {
	Tempo        XmlFloatValue   `xml:"Tempo"`
	TempoEnabled XmlBooleanValue `xml:"TempoEnabled"`
}
