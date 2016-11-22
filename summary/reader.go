package summary

import (
	jenkins "github.com/bndr/gojenkins"
	"time"
)

type Reader interface {
	Read() chan Result
}

type Result struct {
	StatusSummary JobStatusSummary
	Err           error
}

type ReaderOpts struct {
	TimeoutRead time.Duration
	Views []string
}


type readerImpl struct {
	client *jenkins.Jenkins
	opts ReaderOpts
}

func NewReader(client *jenkins.Jenkins, opts ReaderOpts) Reader {
	return &readerImpl{client: client, opts: opts}
}

func (reader *readerImpl) Read() chan Result {
	out := make(chan Result)
	go func() {
		for {
			jobsView, err := reader.getJobsView()
			if err != nil {
				out <- Result{Err: err}
			} else {
				out <- Result{StatusSummary: getJobStatusSummary(jobsView)}
			}
			time.Sleep(reader.opts.TimeoutRead * time.Millisecond)
		}
	}()
	return out
}

type JobsView map[string] []string

func (reader *readerImpl) getJobsView() (JobsView, error) {
	// if view list is not presented
	if len(reader.opts.Views) == 0 {
		// read all jobs
		jobs, err := reader.client.GetAllJobs()
		if err != nil {
			return JobsView{}, err
		}

		colors := []string{}
		for _, job := range jobs {
			colors = append(colors, job.Raw.Color)
		}

		return JobsView{ "All": colors }, nil
	}

	jobsView := make(JobsView)
	for _, view := range reader.opts.Views {
		// read jobs in current view
		v, err := reader.client.GetView(view)
		if err != nil {
			return JobsView{}, err
		}

		colors := []string{}
		jobs := v.GetJobs()

		if len(jobs) == 0 {
			return JobsView{}, err
		}

		// fill colors array
		for _, job := range jobs {
			colors = append(colors, job.Color)
		}

		jobsView[view] = colors
	}

	return jobsView, nil
}
