package pipeline

type FrontendPush struct {
	Chan chan interface{} // Used by other modules for simple
}

func NewFrontendPush() *FrontendPush {
	return &FrontendPush{
		Chan: make(chan interface{}, 10000),
	}
}
