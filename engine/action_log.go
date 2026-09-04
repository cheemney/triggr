package engine

import "log"

type LogAction struct {
	Message string
}

func (a *LogAction) Execute() error {
	log.Println(a.Message)
	return nil
}
