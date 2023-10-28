package pipeline

type FilesForProcessor struct {
	Chan chan *FilesForProcessorMsg
}

type FilesForProcessorMsg struct {
	AbsPath string
}

func NewFilesForProcessor() *FilesForProcessor {
	return &FilesForProcessor{
		Chan: make(chan *FilesForProcessorMsg, 10000),
	}
}
