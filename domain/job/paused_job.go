package job

import "errors"

var ErrPausePaused = errors.New("a job is already paused")

type pausedJob struct {
	job
}

func (j *pausedJob) Process() error {
	j.state = StateProcessing

	return nil
}

func (*pausedJob) Pause() error {
	return ErrPausePaused
}
