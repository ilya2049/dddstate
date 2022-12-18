package job

type State string

const (
	StatePaused     State = "paused"
	StateProcessing State = "processing"
)
