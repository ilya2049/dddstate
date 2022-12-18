package job

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) NewJob(
	id ID,
	name Name,
	state State,
) Job {
	if state == StatePaused {
		return f.newPausedJob(id, name)
	}

	return f.newProcessingJob(id, name)
}

func (f *Factory) newPausedJob(
	id ID,
	name Name,
) Job {
	return &pausedJob{
		job: job{
			id:    id,
			name:  name,
			state: StatePaused,
		},
	}
}

func (f *Factory) newProcessingJob(
	id ID,
	name Name,
) Job {
	return &processingJob{
		job: job{
			id:    id,
			name:  name,
			state: StateProcessing,
		},
	}
}
