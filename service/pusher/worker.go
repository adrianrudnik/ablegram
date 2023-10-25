package pusher

//
//type WorkerPool struct {
//	paths       []string
//	ConsoleChan chan *pipeline.FrontendConsoleMessage
//	limiter     *rate.Limiter
//}
//
//func NewWorkerPool(workerCount int, filesChan chan<- *pipeline.FilesForProcessorMsg) *WorkerPool {
//	return &WorkerPool{
//		ConsoleChan: make(chan *pipeline.FrontendConsoleMessage, 10000),
//		limiter:     rate.NewLimiter(rate.Every(200*time.Millisecond), 1),
//	}
//}
//
//func (p *WorkerPool) Run() {
//	Logger.Info().Msg("Starting frontend push worker")
//
//	// Spool up workers first
//	go p.doWork()
//}
//
//func (p *WorkerPool) doWork() {
//	// The web push must be rate limited to specific intervals, to give the UI time to breathe.
//	// So instead of working single messages as they come in, we process all available in batches
//	// in timed intervals.
//	for {
//		err := p.limiter.Wait(context.Background())
//		if err != nil {
//			Logger.Error().Err(err).Msg("Limiter could not be created")
//			break
//		}
//	}
//
//	// Process status messages
//	batch := make([][]byte, 0)
//
//	for msg := range p.ConsoleChan {
//		b, err := json.Marshal(msg)
//		if err != nil {
//			Logger.Warn().Err(err).Msg("Could not create json for status message")
//			continue
//		}
//
//		batch = append(batch, b)
//	}
//
//	// @todo push the batch towards the websocket
//}
