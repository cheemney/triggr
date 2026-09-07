package config

type Program struct {
	Rules []RuleDecl
}

type RuleDecl struct {
	Name    string
	Trigger Spec
	Action  Spec
}

// Spec is a generic "kind + args" shape shared by triggers and
// actions, e.g. `interval 2s` or `shell "echo hi"`.
type Spec struct {
	Kind string
	Args []string
}
