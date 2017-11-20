package strings

import (
	"fmt"
	"github.com/noelniles/flow/constants"
	"github.com/noelniles/flow/processing"
)

type printer struct {
	processing.Process
	In  chan string
	Out chan string
}

func NewPrinter() *printer {
	return &printer{In: make(chan string, constants.BUFSIZE)}

}

func (proc *printer) Run() {
	for s := range proc.In {
		fmt.Println(s)
	}
}
