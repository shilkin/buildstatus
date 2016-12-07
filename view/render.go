package view

import "github.com/shilkin/buildstatus/status"

type Render interface {
	Render(status.Result) error
}
