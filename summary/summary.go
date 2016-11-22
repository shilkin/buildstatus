package summary

type ViewStatus map[string]JobStatus

func (summary ViewStatus) Add(view string, status JobStatus) {
	current, ok := summary[view]
	if !ok {
		summary[view] = status
		return
	}
	summary[view] = mergeStatus(current, status)
}

func mergeStatus(current JobStatus, new JobStatus) JobStatus {
	if current > new {
		return current
	}
	return new
}
