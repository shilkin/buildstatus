package summary

type JobStatus int

const (
	FAILED JobStatus = iota
	SUCCESS
	INPROGRESS
)

var statusByColor = map[string]JobStatus{
	"blue":       SUCCESS,
	"red":        FAILED,
	"grey":       SUCCESS,
	"blue_anime": INPROGRESS,
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
