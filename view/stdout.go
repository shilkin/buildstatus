package view

import (
	"github.com/shilkin/buildstatus/summary"
	"log"
)

type stdoutRender struct {
}

func (r *stdoutRender) Render(summary summary.ViewStatus) error {
	log.Printf("%#v", summary)
	return nil
}

func NewStdoutRender() Render {
	return &stdoutRender{}
}
