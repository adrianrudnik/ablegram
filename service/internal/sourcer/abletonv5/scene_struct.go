package abletonv5

type SceneDocument struct {
	HasBase
	HasFileReference
	HasUserInfoText
	HasColor
	HasTempoWithToggle
}

func NewSceneDocument() *SceneDocument {
	return &SceneDocument{
		HasBase:          NewHasBase(AbletonScene),
		HasFileReference: NewHasFileReference(),
		HasUserInfoText:  NewHasUserInfoText(),
		HasColor:         NewHasColor(),
		HasTempoWithToggle: HasTempoWithToggle{
			Tempo:        0,
			TempoEnabled: false,
		},
	}

}
