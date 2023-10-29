package webservice

type OpenInput struct {
	Path string `json:"path" binding:"required"`
}
