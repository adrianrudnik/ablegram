package pusher

type ProcessingStatusPush struct {
	Type   string `json:"type"`
	Status int64  `json:"routines"`
}

func NewProcessingStatusPush(status int64) *ProcessingStatusPush {
	return &ProcessingStatusPush{
		Type:   "processing_status",
		Status: status,
	}
}
