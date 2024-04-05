package shell

import (
	"fmt"
	"strings"
)

type EnvStatus struct {
	Added    []string
	Replaced []string
}

type Status struct {
	Env EnvStatus
}

func (s Status) Save() {

}

func (s Status) Short() string {
	return fmt.Sprintf("+%d <>%d", len(s.Env.Added), len(s.Env.Replaced))
}

func (s Status) String() string {
	var status []string

	for _, added := range s.Env.Added {
		status = append(status, fmt.Sprintf("+%s", added))
	}

	for _, replaced := range s.Env.Replaced {
		status = append(status, fmt.Sprintf("<>%s", replaced))
	}

	return strings.Join(status, " ")
}
