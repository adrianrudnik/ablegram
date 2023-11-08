package workload

import (
	"fmt"
)

type DocumentPayload struct {
	Id       string
	Document interface{}
}

func NewDocumentPayload(id string, doc interface{}) *DocumentPayload {
	return &DocumentPayload{
		Id:       fmt.Sprintf("%s", id),
		Document: doc,
	}
}
