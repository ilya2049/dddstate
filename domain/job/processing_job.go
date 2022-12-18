package job

import "errors"

var ErrProcessProcessing = errors.New("a job is already processing")

type processingJob struct {
	job
}

func (*processingJob) Process() error {
	return ErrProcessProcessing
}

func (j *processingJob) Pause() error {
	j.state = StateProcessing

	return nil
}
