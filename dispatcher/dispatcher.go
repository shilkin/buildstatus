package dispatcher

import (
	"github.com/shilkin/buildstatus/status"
	"github.com/shilkin/buildstatus/view"
)

type Dispatcher interface {
	Run() error
}

type dispatcherImpl struct {
	reader status.Reader
	render view.Render
}

func (d *dispatcherImpl) Run() (err error) {
	for status := range d.reader.Read() {
		if status.Err != nil {
			continue
		}
		_ = d.render.Render(status)
	}
	return
}

func NewDispatcher(reader status.Reader, render view.Render) Dispatcher {
	return &dispatcherImpl{reader: reader, render: render}
}
