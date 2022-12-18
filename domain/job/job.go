package job

type Job interface {
	ID() ID
	Name() Name
	State() State
	Process() error
	Pause() error
}

type job struct {
	id    ID
	name  Name
	state State
}

func (j *job) ID() ID {
	return j.id
}

func (j *job) State() State {
	return j.state
}

func (j *job) Name() Name {
	return j.name
}
