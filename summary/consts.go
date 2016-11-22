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
