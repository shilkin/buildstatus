package view

import "github.com/shilkin/buildstatus/summary"

type Render interface {
	Render(summary.JobStatusSummary) error
}
