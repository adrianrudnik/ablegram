package pusher

import (
	"crypto/md5"
	"fmt"
	"path/filepath"
)

type FileStatusPush struct {
	Type     string `json:"type"`
	ID       string `json:"id"`
	AbsPath  string `json:"path"`
	Folder   string `json:"folder"`
	Filename string `json:"filename"`
	Status   string `json:"status"`
	Remark   string `json:"remark,omitempty"`
}

func NewFileStatusPush(path string, status string, remark string) *FileStatusPush {
	id := md5.Sum([]byte(path))

	p := &FileStatusPush{
		Type:     "file_status",
		ID:       fmt.Sprintf("%x", id),
		AbsPath:  path,
		Folder:   filepath.Dir(path),
		Filename: filepath.Base(path),
		Status:   status,
	}

	if remark != "" {
		p.Remark = remark
	}

	return p
}

type MetricUpdatePush struct {
	Type   string            `json:"type"`
	Values map[string]uint64 `json:"values"`
}

func NewMetricUpdatePush(values map[string]uint64) *MetricUpdatePush {
	return &MetricUpdatePush{
		Type:   "metric_update",
		Values: values,
	}
}

type ProcessingStatusPush struct {
	Type   string `json:"type"`
	Status bool   `json:"status"`
}

func NewProcessingStatusPush(status bool) *ProcessingStatusPush {
	return &ProcessingStatusPush{
		Type:   "processing_status",
		Status: status,
	}
}
