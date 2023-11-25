package pusher

type TagUpdatePush struct {
	Type   string            `json:"type"`
	Values map[string]uint64 `json:"tags"`
}

func NewTagUpdatePush(values map[string]uint64) *TagUpdatePush {
	return &TagUpdatePush{
		Type:   "tag_update",
		Values: values,
	}
}
