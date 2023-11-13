package abletonv5

import (
	"crypto/md5"
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/util"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// HasBase is the minimum struct to comply to for the indexer
type HasBase struct {
	T            string   `json:"type"`
	Tags         []string `json:"tags,omitempty"`
	DisplayName  string   `json:"displayName,omitempty"`
	PathAbsolute string   `json:"pathAbsolute,omitempty"`
	PathFolder   string   `json:"pathFolder,omitempty"`
	Filename     string   `json:"filename,omitempty"`
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

func (b *HasBase) EngraveTags(tb *tagger.TagBucket) {
	b.Tags = tb.Engrave([]string{b.PathAbsolute})
}

func NewHasBase(t string) HasBase {
	return HasBase{T: t}
}

func (b *HasBase) LoadFileReference(path string, tb *tagger.TagBucket) {
	b.PathAbsolute = path
	b.PathFolder = filepath.Dir(path)
	b.Filename = filepath.Base(path)

	// Determine the overall location of the file
	homeDir, err := os.UserHomeDir()
	if err == nil {
		if strings.HasPrefix(path, homeDir) {
			tb.Add("file:location=inside-user-home")
		} else {
			tb.Add("file:location=outside-user-home")
		}
	}

	// Backups can be caught by a path pattern like
	// ".../samples/Backup/MIDI Effect Arpeggiator [2023-11-06 163730].als"
	found, err := regexp.MatchString(`Backup[/\\](.*)\[\d{4}-\d{2}-\d{2} \d{6}]`, path)
	if err == nil && found {
		tb.Add("file:location=ableton-backup")
	}

	if util.PathContainsFolder(path, "Trash") || util.PathContainsFolder(path, "$Recycle.Bin") {
		tb.Add("file:location=trash")
	}

	if util.PathContainsFolder(path, "pCloudDrive") {
		tb.Add("file:location=p-cloud")
	}

	if util.PathContainsFolder(path, "Live Recordings") {
		tb.Add("file:location=ableton-live-recording")
	}

	if util.PathContainsFolder(path, "Factory Packs") {
		tb.Add("file:location=ableton-factory-pack")
	}

	if util.PathContainsFolder(path, "Cloud Manager") {
		tb.Add("file:location=ableton-cloud-manager")
	}

	if util.PathContainsFolder(path, "User Library") {
		tb.Add("file:location=ableton-user-library")
	}
}

type HasUserName struct {
	UserName string `json:"userName,omitempty"`
}

func NewHasUserName() HasUserName {
	return HasUserName{}
}

func (u *HasUserName) LoadUserName(v string, t *tagger.TagBucket) {
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

func (f *HasTrackUserNames) LoadTrackUserNames(v *XmlTrackNameNode, t *tagger.TagBucket) {
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

func (c *HasColor) LoadColor(v int16, t *tagger.TagBucket) {
	c.Color = v

	t.Add(fmt.Sprintf("color:ableton=%d", v))
}

// HasUserInfoText represents an element that can be annotated by the user, also known as "Info Text".
type HasUserInfoText struct {
	Annotation string `json:"annotation,omitempty"`
}

func (a *HasUserInfoText) LoadUserInfoText(v string, t *tagger.TagBucket) {
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

func (t *HasTempoWithToggle) LoadTempoWithToggle(v *XmlTempoWithToggleNode, tags *tagger.TagBucket) {
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

type HasIsExpandedOption struct {
	IsExpanded bool `json:"isExpanded,omitempty"`
}

func NewHasDeviceIsExpanded() HasIsExpandedOption {
	return HasIsExpandedOption{IsExpanded: false}
}

func (h *HasIsExpandedOption) LoadOptionExpanded(v bool, tags *tagger.TagBucket) {
	h.IsExpanded = v

	if v {
		tags.Add("ableton-option:expanded=true")
	} else {
		tags.Add("ableton-option:expanded=false")
	}
}

type HasIsFoldedOption struct {
	IsFolded bool `json:"isFolded,omitempty"`
}

func NewHasDeviceIsFolded() HasIsFoldedOption {
	return HasIsFoldedOption{IsFolded: false}
}

func (h *HasIsFoldedOption) LoadOptionFolded(v bool, tags *tagger.TagBucket) {
	h.IsFolded = v

	if v {
		tags.Add("ableton-option:folded=true")
	} else {
		tags.Add("ableton-option:folded=false")
	}
}

type HasIsFrozenOption struct {
	IsFrozen bool `json:"isFrozen,omitempty"`
}

func NewHasTrackIsFrozen() HasIsFrozenOption {
	return HasIsFrozenOption{IsFrozen: false}
}

func (h *HasIsFrozenOption) LoadIsFrozenOption(v bool, tags *tagger.TagBucket) {
	h.IsFrozen = v

	if v {
		tags.Add("ableton-option:frozen=true")
	} else {
		tags.Add("ableton-option:frozen=false")
	}
}

type HasScaleInformation struct {
	ScaleRootNote string `json:"scaleRootNote,omitempty"`
	ScaleName     string `json:"scaleName,omitempty"`
}

func NewHasScaleInformation() HasScaleInformation {
	return HasScaleInformation{}
}

func (h *HasScaleInformation) LoadScaleInformation(v *XmlScaleInformationValue, tags *tagger.TagBucket) {
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

func (h *HasTimeSignature) LoadTimeSignature(v *XmlRemoteableTimeSignature, tags *tagger.TagBucket) {
	h.TimeSignature = fmt.Sprintf("%d/%d", v.Numerator.Value, v.Denominator.Value)

	tags.Add(fmt.Sprintf("time-signature:name=%s", h.TimeSignature))
	tags.Add(fmt.Sprintf("time-signature:numerator=%d", v.Numerator.Value))
	tags.Add(fmt.Sprintf("time-signature:denominator=%d", v.Denominator.Value))
}
