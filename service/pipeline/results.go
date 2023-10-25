package pipeline

import (
	"crypto/md5"
	"fmt"
)

type DocumentsToIndex struct {
	Channel chan *DocumentToIndexMsg
}

type DocumentToIndexMsg struct {
	Id       string
	Document interface{}
}

func NewDocumentToIndexMsg(id string, doc interface{}) *DocumentToIndexMsg {
	return &DocumentToIndexMsg{
		Id:       fmt.Sprintf("%x", md5.Sum([]byte(id))),
		Document: doc,
	}
}

func NewDocumentsToIndex() *DocumentsToIndex {
	return &DocumentsToIndex{
		Channel: make(chan *DocumentToIndexMsg, 10000),
	}
}
