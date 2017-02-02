package view

import (
	"github.com/shilkin/buildstatus/status"
	"errors"
	"fmt"
)

type Render interface {
	Render(status.Result) error
}

func NewRender(renderType string, opts interface{}) (Render, error) {
	switch renderType {
	case "stdout":
		return NewStdoutRender(), nil
	case "trafficlight":
		o, ok := opts.(RaspberryOpts)
		if !ok {
			return nil, errors.New("invalid options type")
		}
		return NewRaspberryRender(o)
	}

	return nil, fmt.Errorf("invalid render type '%s'", renderType)
}
