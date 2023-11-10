package abletonv5

import (
	"fmt"
	"math"
	"strings"
)

type XmlStringValue struct {
	Value string `xml:"Value,attr"`
}

type XmlIntValue struct {
	Value int64 `xml:"Value,attr"`
}

type XmlFloatValue struct {
	Value float64 `xml:"Value,attr"`
}

type XmlBooleanValue struct {
	Value bool `xml:"Value,attr"`
}

type XmlColorValue struct {
	Value int16 `xml:"Value,attr"`
}

type XmlMidiKey struct {
	Value int `xml:"Value,attr"`
}

func (x *XmlMidiKey) HumanReadable(includeOctave bool) string {
	// @see https://www.zem-college.de/midi/mc_taben.htm
	// @see https://computermusicresource.com/midikeys.html
	octave := int(math.Floor(float64(x.Value/12))) - 2
	key := x.Value - ((octave + 2) * 12)

	var tmpl string
	switch key {
	case 0:
		tmpl = "C%d"
	case 1:
		tmpl = "C#%d"
	case 2:
		tmpl = "D%d"
	case 3:
		tmpl = "D#%d"
	case 4:
		tmpl = "E%d"
	case 5:
		tmpl = "F%d"
	case 6:
		tmpl = "F#%d"
	case 7:
		tmpl = "G%d"
	case 8:
		tmpl = "G#%d"
	case 9:
		tmpl = "A%d"
	case 10:
		tmpl = "A#%d"
	case 11:
		tmpl = "B%d"
	default:
		tmpl = "unknown"
	}

	if includeOctave {
		return fmt.Sprintf(tmpl, octave)
	} else {
		return strings.Replace(tmpl, "%d", "", 1)
	}
}
