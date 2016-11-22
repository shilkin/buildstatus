package view

import (
	"github.com/shilkin/buildstatus/summary"
	"log"
)

type stdoutRender struct {

}

func (r *stdoutRender) Render(summary summary.JobStatusSummary) (err error) {
	if len(summary) == 0 {
		log.Printf("jobs are in progress...")
		return
	}
	log.Printf("%#v", summary)
	return
}

func NewStdoutRender() Render {
	return &stdoutRender{}
}
