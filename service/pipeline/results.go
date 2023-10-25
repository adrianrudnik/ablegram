package pipeline

type ResultsToIndex struct {
	Channel chan *ResultToIndexMsg
}

type ResultToIndexMsg struct {
}

func NewResultsToIndex() *ResultsToIndex {
	return &ResultsToIndex{
		Channel: make(chan *ResultToIndexMsg, 10000),
	}
}
