package summary

type JobStatusSummary map[string] JobStatus

func (summary JobStatusSummary) Add(view string, status JobStatus) {
	current, ok := summary[view]
	if !ok {
		current = SUCCESS
		summary[view] = current
	}
	summary[view] = mergeStatus(current, status)
}

func mergeStatus(current JobStatus, new JobStatus) JobStatus {
	if current == FAILED || new == FAILED {
		return FAILED
	}
	if current == INPROGRESS || new == INPROGRESS {
		return INPROGRESS
	}
	return SUCCESS
}
