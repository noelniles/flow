package strings

import (
	"fmt"
	"github.com/noelniles/flow/constants"
	"time"
)

type time_sayer struct {
	Out chan string
}

func NewTimeSayer() *time_sayer {
	return &time_sayer{make(chan string, constants.BUFSIZE)}
}

func (proc *time_sayer) Run() {
	defer close(proc.Out)
	for {
		fmt.Println(time.Now())
	}
}
