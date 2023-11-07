package abletonv5

import (
	"crypto/md5"
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/util"
	"math"
	"path/filepath"
)

// HasBase is the minimum struct to comply to for the indexer
type HasBase struct {
	T           string   `json:"type"`
	Tags        []string `json:"tags,omitempty"`
	DisplayName string   `json:"displayName,omitempty"`
}

func (b *HasBase) Type() string {
	return b.T
}

func (b *HasBase) LoadDisplayName(parts []string) {
	b.DisplayName = util.Namelize(parts)
}

func (b *HasBase) GetAutoId() string {
	id := idGenerator.Add(1)
	typedId := fmt.Sprintf("%s_%d", b.T, id)

	return fmt.Sprintf("%x", md5.Sum([]byte(typedId)))
}

func (b *HasBase) EngraveTags(t *tagger.Tagger) {
	b.Tags = t.Engrave()
}

func NewHasBase(t string) HasBase {
	return HasBase{T: t}
}

// HasFileReference represents a link to a file that contained the element
type HasFileReference struct {
	PathAbsolute string `json:"pathAbsolute,omitempty"`
	PathFolder   string `json:"pathFolder,omitempty"`
	Filename     string `json:"filename,omitempty"`
}

func (r *HasFileReference) LoadFileReference(path string, t *tagger.Tagger) {
	r.PathAbsolute = path
	r.PathFolder = filepath.Dir(path)
	r.Filename = filepath.Base(path)
}

func NewHasFileReference() HasFileReference {
	return HasFileReference{}
}

type HasUserName struct {
	UserName string `json:"userName,omitempty"`
}

func NewHasUserName() HasUserName {
	return HasUserName{}
}

func (u *HasUserName) LoadUserName(v string, t *tagger.Tagger) {
	m, empty := util.EvaluateUserInput(v)

	if !empty {
		u.UserName = m
	}

	if t != nil {
		if !empty {
			t.Add("user:name=true")
		} else {
			t.Add("user:name=false")
		}
	}
}

// HasTrackUserNames represents an element that can be named by the user.
// It contains of many fields that are a composition of many nameable, seen in tracks.
type HasTrackUserNames struct {
	HasUserName
	HasUserInfoText

	EffectiveName          string `json:"effectiveName,omitempty"`
	MemorizedFirstClipName string `json:"memorizedFirstClipName,omitempty"`
}

func NewHasTrackUserNames() HasTrackUserNames {
	return HasTrackUserNames{
		HasUserName:     NewHasUserName(),
		HasUserInfoText: NewHasUserInfoText(),
	}
}

func (f *HasTrackUserNames) LoadTrackUserNames(v *XmlTrackNameNode, t *tagger.Tagger) {
	f.LoadUserName(v.Name.UserName.Value, t)
	f.LoadUserInfoText(v.Name.Annotation.Value, t)

	f.EffectiveName = v.Name.EffectiveName.Value
	f.MemorizedFirstClipName = v.Name.MemorizedFirstClipName.Value
}

// HasColor represents an element that can be colored by the user.
type HasColor struct {
	Color int16 `json:"color,omitempty"`
}

func NewHasColor() HasColor {
	return HasColor{}
}

func (c *HasColor) LoadColor(v int16, t *tagger.Tagger) {
	c.Color = v

	t.Add(fmt.Sprintf("color:ableton=%d", v))
}

// HasUserInfoText represents an element that can be annotated by the user, also known as "Info Text".
type HasUserInfoText struct {
	Annotation string `json:"annotation,omitempty"`
}

func (a *HasUserInfoText) LoadUserInfoText(v string, t *tagger.Tagger) {
	m, empty := util.EvaluateUserInput(v)

	if !empty {
		a.Annotation = m
	}

	if t != nil {
		if !empty {
			t.Add("user:memo=true")
		} else {
			t.Add("user:memo=false")
		}
	}
}

func NewHasUserInfoText() HasUserInfoText {
	return HasUserInfoText{}
}

type HasTempo struct {
	Tempo float64 `json:"tempo,omitempty"`
}

type HasTempoWithToggle struct {
	Tempo        float64 `json:"tempo,omitempty"`
	TempoEnabled bool    `json:"tempoEnabled,omitempty"`
}

func NewHasTempoWithToggle() HasTempoWithToggle {
	return HasTempoWithToggle{
		Tempo:        0,
		TempoEnabled: false,
	}
}

func (t *HasTempoWithToggle) LoadTempoWithToggle(v *XmlTempoWithToggleNode, tags *tagger.Tagger) {
	t.Tempo = v.Tempo.Value
	t.TempoEnabled = v.TempoEnabled.Value

	if (v.Tempo.Value > 0) && v.TempoEnabled.Value {
		if math.Trunc(v.Tempo.Value) == v.Tempo.Value {
			// If we have a rounded tempo, we just need to add one tag
			tags.Add(fmt.Sprintf("bpm:%d", int(math.Round(v.Tempo.Value))))
		} else {
			// Otherwise it's a weird file where the tempo is a fraction, like in some XmlRoot delivered ALS files.
			// We just add both rounded values to the tags
			tags.Add(fmt.Sprintf("bpm:%d", int(math.Floor(v.Tempo.Value))))
			tags.Add(fmt.Sprintf("bpm:%d", int(math.Ceil(v.Tempo.Value))))
		}
	}
}

type HasDeviceIsExpanded struct {
	IsExpanded bool `json:"isExpanded,omitempty"`
}

func NewHasDeviceIsExpanded() HasDeviceIsExpanded {
	return HasDeviceIsExpanded{IsExpanded: false}
}

func (h *HasDeviceIsExpanded) LoadDeviceIsExpanded(v bool, tags *tagger.Tagger) {
	h.IsExpanded = v

	if v {
		tags.Add("ableton-device:expanded=true")
	} else {
		tags.Add("ableton-device:expanded=false")
	}
}

type HasDeviceIsFolded struct {
	IsFolded bool `json:"isFolded,omitempty"`
}

func NewHasDeviceIsFolded() HasDeviceIsFolded {
	return HasDeviceIsFolded{IsFolded: false}
}

func (h *HasDeviceIsFolded) LoadDeviceIsFolded(v bool, tags *tagger.Tagger) {
	h.IsFolded = v

	if v {
		tags.Add("ableton-device:folded=true")
	} else {
		tags.Add("ableton-device:folded=false")
	}
}

type HasTrackIsFrozen struct {
	IsFrozen bool `json:"isFrozen,omitempty"`
}

func NewHasTrackIsFrozen() HasTrackIsFrozen {
	return HasTrackIsFrozen{IsFrozen: false}
}

func (h *HasTrackIsFrozen) LoadTrackIsFrozen(v bool, tags *tagger.Tagger) {
	h.IsFrozen = v

	if v {
		tags.Add("ableton-track:frozen=true")
	} else {
		tags.Add("ableton-track:frozen=false")
	}
}

type HasScaleInformation struct {
	ScaleRootNote string `json:"scaleRootNote,omitempty"`
	ScaleName     string `json:"scaleName,omitempty"`
}

func NewHasScaleInformation() HasScaleInformation {
	return HasScaleInformation{}
}

func (h *HasScaleInformation) LoadScaleInformation(v *XmlScaleInformationValue, tags *tagger.Tagger) {
	h.ScaleRootNote = v.HumanizeRootNote()
	h.ScaleName = v.HumanizeName()

	if h.ScaleRootNote != "" {
		tags.Add(fmt.Sprintf("scale:root-note=%s", h.ScaleRootNote))
	}

	if h.ScaleName != "" {
		tags.Add(fmt.Sprintf("scale:name=%s", h.ScaleName))
	}
}

type HasTimeSignature struct {
	TimeSignature string `json:"timeSignature,omitempty"`
}

func NewHasTimeSignature() HasTimeSignature {
	return HasTimeSignature{}
}

func (h *HasTimeSignature) LoadTimeSignature(v *XmlRemoteableTimeSignature, tags *tagger.Tagger) {
	h.TimeSignature = fmt.Sprintf("%d/%d", v.Numerator.Value, v.Denominator.Value)

	tags.Add(fmt.Sprintf("time-signature=%s", h.TimeSignature))
}
