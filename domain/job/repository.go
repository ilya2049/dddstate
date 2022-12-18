package job

import "context"

type Repository interface {
	Get(context.Context, ID) (Job, error)
	Update(context.Context, Job) error
}
