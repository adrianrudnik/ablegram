package pipeline

type FrontendPush struct {
	Channel chan interface{} // Used by other modules for simple
}

func NewFrontendPush() *FrontendPush {
	return &FrontendPush{
		Channel: make(chan interface{}, 10000),
	}
}
