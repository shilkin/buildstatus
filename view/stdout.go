package view

import (
	"github.com/shilkin/buildstatus/status"
	"log"
)

type stdoutRender struct {
}

func (r *stdoutRender) Render(summary status.Result) error {
	log.Printf("%#v (err = '%v')", summary.StatusSummary, summary.Err)
	return nil
}

func NewStdoutRender() Render {
	return &stdoutRender{}
}
