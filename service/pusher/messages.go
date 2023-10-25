package pusher

type FileStatusPush struct {
	Path   string `json:"path"`
	Status string `json:"status"`
	Remark string `json:"remark,omitempty"`
}
