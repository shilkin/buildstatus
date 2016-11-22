package summary

var (
	DEPLOY_PREFIX        = []string{"DEPLOY"}
	BUILD_LOCAL_PREFIX   = []string{"LOCAL"}
	BUILD_FOREIGN_PREFIX = []string{"FOREIGN"}
	TEST_PREFIX          = []string{"AT", "ACCEPTANCE", "SMOKE"}
	DOCKER_PREFIX        = []string{"DOCKER"}
	ENV_PREFIX           = []string{"ENV"}
	RUN_PREFIX           = []string{"RUN"}
)

type jobType int

const (
	DEPLOY jobType = iota
	BUILD_LOCAL
	BUILD_FOREIGN
	TEST
	DOCKER
	ENV
	RUN
)


