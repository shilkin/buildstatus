package view

import (
	"github.com/shilkin/buildstatus/status"
	"log"
)

type stdoutRender struct {
}

func (r *stdoutRender) Render(summary status.ViewStatus) error {
	log.Printf("%#v", summary)
	return nil
}

func NewStdoutRender() Render {
	return &stdoutRender{}
}
