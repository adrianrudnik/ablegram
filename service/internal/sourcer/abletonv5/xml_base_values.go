package abletonv5

import (
	"fmt"
	"math"
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

func (x *XmlMidiKey) HumanReadable() string {
	// @see https://www.zem-college.de/midi/mc_taben.htm
	// @see https://computermusicresource.com/midikeys.html
	octave := int(math.Floor(float64(x.Value/12))) - 2
	key := x.Value - ((octave + 2) * 12)

	switch key {
	case 0:
		return fmt.Sprintf("C%d", octave)
	case 1:
		return fmt.Sprintf("C#%d", octave)
	case 2:
		return fmt.Sprintf("D%d", octave)
	case 3:
		return fmt.Sprintf("D#%d", octave)
	case 4:
		return fmt.Sprintf("E%d", octave)
	case 5:
		return fmt.Sprintf("F%d", octave)
	case 6:
		return fmt.Sprintf("F#%d", octave)
	case 7:
		return fmt.Sprintf("G%d", octave)
	case 8:
		return fmt.Sprintf("G#%d", octave)
	case 9:
		return fmt.Sprintf("A%d", octave)
	case 10:
		return fmt.Sprintf("A#%d", octave)
	case 11:
		return fmt.Sprintf("B%d", octave)
	}

	return "unknown"
}
