package status

type ViewStatus map[string]JobStatus

func (s ViewStatus) Add(view string, status JobStatus) {
	current, ok := s[view]
	if !ok {
		s[view] = status
		return
	}
	s[view] = mergeStatus(current, status)
}

func mergeStatus(current JobStatus, new JobStatus) JobStatus {
	if current > new {
		return current
	}
	return new
}
