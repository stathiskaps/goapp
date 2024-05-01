package watcher

type Counter struct {
	Iteration int    `json:"iteration"`
	Message   string `json:"message"`
}

type CounterReset struct{}
