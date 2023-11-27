package workload

type FilePayload struct {
	AbsPath string
}

func NewFilePayload(absPath string) *FilePayload {
	return &FilePayload{
		AbsPath: absPath,
	}
}
