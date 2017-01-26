package view

import (
	"github.com/davecheney/gpio"
	"github.com/shilkin/buildstatus/status"
	"sync"
	"time"
)

type RaspberryOpts struct {
	GpioGreen  int
	GpioYellow int
	GpioRed    int
}

type modeType int

const (
	green modeType = iota
	yellow
	red
	blink
)

const (
	blinkTimeout = 500
)

type raspberryRender struct {
	pinGreen    gpio.Pin
	pinYellow   gpio.Pin
	pinRed      gpio.Pin
	stopBlinkCh chan struct{}
	isBlinking  bool
	mutex       sync.RWMutex
}

func NewRaspberryRender(opts RaspberryOpts) (render Render, err error) {
	pinGreen, err := gpio.OpenPin(opts.GpioGreen, gpio.ModeOutput)
	if err != nil {
		return
	}
	pinYellow, err := gpio.OpenPin(opts.GpioYellow, gpio.ModeOutput)
	if err != nil {
		return
	}
	pinRed, err := gpio.OpenPin(opts.GpioRed, gpio.ModeOutput)
	if err != nil {
		return
	}
	render = &raspberryRender{
		pinGreen:    pinGreen,
		pinYellow:   pinYellow,
		pinRed:      pinRed,
		stopBlinkCh: make(chan struct{}),
	}
	return
}

func (r *raspberryRender) Render(summary status.Result) (err error) {
	if summary.Err != nil {
		r.yellow()
		return
	}

	// merge all statuses by priority
	current := status.SUCCESS
	for _, s := range summary.StatusSummary {
		if s > current {
			current = s
		}
	}

	// choose mode according to status
	switch current {
	case status.SUCCESS:
		r.green()
	case status.FAILED:
		r.red()
	case status.INPROGRESS:
		r.blink()
	}

	return
}

func (r *raspberryRender) green() {
	r.clear()
	r.pinGreen.Set()
}

func (r *raspberryRender) yellow() {
	r.clear()
	r.pinYellow.Set()
}

func (r *raspberryRender) red() {
	r.clear()
	r.pinRed.Set()
}

func (r *raspberryRender) blink() {
	r.clear()
	// start blink goroutine
	r.setBlinkingStatus(true)
	go func() {
		for {
			select {
			case <-r.stopBlinkCh:
				r.setBlinkingStatus(false)
				break
			default:
				break
			}
			r.pinYellow.Set()
			time.Sleep(blinkTimeout * time.Millisecond)
			r.pinYellow.Clear()
			time.Sleep(blinkTimeout * time.Millisecond)
		}
	}()
}

func (r *raspberryRender) clear() {
	r.pinGreen.Clear()
	r.pinRed.Clear()
	r.pinYellow.Clear()
	r.stopBlink()
}

func (r *raspberryRender) stopBlink() {
	// stop blink goroutine
	r.mutex.RLock()
	if r.isBlinking {
		r.stopBlinkCh <- struct{}{}
	}
	r.mutex.RUnlock()
}

func (r *raspberryRender) setBlinkingStatus(isBlinking bool) {
	r.mutex.Lock()
	r.isBlinking = isBlinking
	r.mutex.Unlock()
}
