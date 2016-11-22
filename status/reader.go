package status

import (
	jenkins "github.com/bndr/gojenkins"
	"time"
)

type Reader interface {
	Read() chan Result
}

type Result struct {
	StatusSummary ViewStatus
	Err           error
}

type ReaderOpts struct {
	TimeoutRead time.Duration
	Views       []string
}

type jenkinsReader struct {
	client *jenkins.Jenkins
	opts   ReaderOpts
}

func NewReader(client *jenkins.Jenkins, opts ReaderOpts) Reader {
	return &jenkinsReader{client: client, opts: opts}
}

func (r *jenkinsReader) Read() chan Result {
	out := make(chan Result)
	go func() {
		for {
			jobsSummary, err := r.getViewsStatus()
			if err != nil {
				out <- Result{Err: err}
			} else {
				out <- Result{StatusSummary: jobsSummary}
			}
			time.Sleep(r.opts.TimeoutRead * time.Millisecond)
		}
	}()
	return out
}

func (r *jenkinsReader) getViewsStatus() (result ViewStatus, err error) {
	result = make(ViewStatus)

	views, err := r.getViews()
	if err != nil {
		return
	}

ViewLoop:
	for _, view := range views {
		var v *jenkins.View
		v, err = r.client.GetView(view)
		if err != nil {
			return
		}

		jobs := v.GetJobs()
		if len(jobs) == 0 {
			continue // next view
		}

		for _, job := range jobs {
			status, ok := statusByColor[job.Color]
			if !ok {
				continue // next job
			}
			result.Add(view, status)
			if status == INPROGRESS {
				continue ViewLoop // next view
			}
		}
	}

	return
}

func (reader *jenkinsReader) getViews() (result []string, err error) {
	// if views are set - return them
	if len(reader.opts.Views) != 0 {
		result = reader.opts.Views
		return
	}

	// if view aren't set - return all from server
	views, err := reader.client.GetAllViews()
	if err != nil {
		return
	}
	for _, view := range views {
		result = append(result, view.Raw.Name)
	}
	return
}
