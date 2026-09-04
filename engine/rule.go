package engine

// Trigger is the "if this" half — it blocks until its condition is met.
type Trigger interface {
	Watch() error
}

// Action is the "then that" half — it runs once a Trigger fires.
type Action interface {
	Execute() error
}

type Rule struct {
	Name    string
	Trigger Trigger
	Action  Action
}

func (r *Rule) Run() error {
	if err := r.Trigger.Watch(); err != nil {
		return err
	}
	return r.Action.Execute()
}
