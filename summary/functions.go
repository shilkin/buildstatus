package summary

type JobStatus int

const (
	SUCCESS JobStatus = iota
	FAILED
	INPROGRESS
)

var statusByColor = map[string]JobStatus{
	"blue":       SUCCESS,
	"red":        FAILED,
	"grey":       SUCCESS,
	"blue_anime": INPROGRESS,
	"red_anime":  INPROGRESS,
}

func getJobStatusSummary(jobsView JobsView) (result JobStatusSummary) {
	result = make(JobStatusSummary)
JobsSummaryLoop:
	for view, colors := range jobsView {
		for _, color := range colors {
			status, ok := statusByColor[color]
			if !ok {
				continue JobsSummaryLoop
			}

			result.Add(view, status)

			if status == INPROGRESS {
				continue JobsSummaryLoop
			}
		}
	}
	return
}
