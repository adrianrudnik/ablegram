package abletonsrc

type XmlRemoteableTimeSignatureNode struct {
	TimeSignature XmlRemoteableTimeSignature `xml:"TimeSignature"`
}

type XmlRemoteableTimeSignature struct {
	Numerator   XmlIntValue `xml:"TimeSignatures>RemoteableTimeSignature>Numerator"`
	Denominator XmlIntValue `xml:"TimeSignatures>RemoteableTimeSignature>Denominator"`
	Time        XmlIntValue `xml:"TimeSignatures>RemoteableTimeSignature>Time"`
}
