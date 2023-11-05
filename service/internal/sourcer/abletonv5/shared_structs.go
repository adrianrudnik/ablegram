package abletonv5

import (
	"crypto/md5"
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/util"
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
	b.Tags = t.GetAllAndClear()
}

func NewHasBase(t string) *HasBase {
	return &HasBase{T: t}
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

func NewHasFileReference() *HasFileReference {
	return &HasFileReference{}
}

type HasUserName struct {
	UserName string `json:"userName,omitempty"`
}

func NewHasUserName() *HasUserName {
	return &HasUserName{}
}

func (u *HasUserName) LoadUserName(v string, t *tagger.Tagger) {
	m, empty := util.EvaluateUserInput(v)

	if !empty {
		u.UserName = m
	}

	if t != nil {
		if !empty {
			t.AddSystemTag("info:has-user-name")
		} else {
			t.AddSystemTag("info:no-user-name")
		}
	}
}

// HasTrackUserNames represents an element that can be named by the user.
// It contains of many fields that are a composition of many nameable, seen in tracks.
type HasTrackUserNames struct {
	*HasUserName
	*HasUserInfoText

	EffectiveName          string `json:"effectiveName,omitempty"`
	MemorizedFirstClipName string `json:"memorizedFirstClipName,omitempty"`
}

func NewHasTrackUserNames() *HasTrackUserNames {
	return &HasTrackUserNames{
		HasUserName:     NewHasUserName(),
		HasUserInfoText: NewHasUserInfo(),
	}
}

func (f *HasTrackUserNames) LoadTrackUserNames(v *XmlFullName, t *tagger.Tagger) {
	f.LoadUserName(v.UserName.Value, t)
	f.LoadUserInfoText(v.Annotation.Value, t)

	f.EffectiveName = v.EffectiveName.Value
	f.MemorizedFirstClipName = v.MemorizedFirstClipName.Value
}

// HasColor represents an element that can be colored by the user.
type HasColor struct {
	Color int16 `json:"color,omitempty"`
}

func NewHasColor() *HasColor {
	return &HasColor{}
}

func (c *HasColor) LoadColor(v int16, t *tagger.Tagger) {
	c.Color = v

	t.AddSystemTag(fmt.Sprintf("color:ableton:%d", v))
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
			t.AddSystemTag("info:has-user-memo")
		} else {
			t.AddSystemTag("info:no-user-memo")
		}
	}
}

func NewHasUserInfo() *HasUserInfoText {
	return &HasUserInfoText{}
}
