package pipeline

type FilesForProcessor struct {
	Channel chan *FilesForProcessorMsg
}

type FilesForProcessorMsg struct {
	AbsPath string
}

func NewFilesForProcessor() *FilesForProcessor {
	return &FilesForProcessor{
		Channel: make(chan *FilesForProcessorMsg, 10000),
	}
}
