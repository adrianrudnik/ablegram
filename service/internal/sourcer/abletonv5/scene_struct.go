package abletonv5

type SceneDocument struct {
	HasBase
	HasUserInfoText
	HasColor
	HasTempoWithToggle
}

func NewSceneDocument() *SceneDocument {
	return &SceneDocument{
		HasBase:         NewHasBase(AbletonScene),
		HasUserInfoText: NewHasUserInfoText(),
		HasColor:        NewHasColor(),
		HasTempoWithToggle: HasTempoWithToggle{
			Tempo:        0,
			TempoEnabled: false,
		},
	}

}
