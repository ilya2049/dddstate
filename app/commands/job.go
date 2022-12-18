package commands

import (
	"context"

	"dddstate/domain/job"
)

type JobCommandExecutor struct {
	jobRepository job.Repository
}

func NewJobCommandExecutor(jobRepository job.Repository) *JobCommandExecutor {
	return &JobCommandExecutor{
		jobRepository: jobRepository,
	}
}

func (ex *JobCommandExecutor) ProcessJob(ctx context.Context, jobID job.ID) error {
	aJob, err := ex.jobRepository.Get(ctx, jobID)
	if err != nil {
		return err
	}

	if err := aJob.Process(); err != nil {
		return err
	}

	return ex.jobRepository.Update(ctx, aJob)
}

func (ex *JobCommandExecutor) PauseJob(ctx context.Context, jobID job.ID) error {
	aJob, err := ex.jobRepository.Get(ctx, jobID)
	if err != nil {
		return err
	}

	if err := aJob.Pause(); err != nil {
		return err
	}

	return ex.jobRepository.Update(ctx, aJob)
}
