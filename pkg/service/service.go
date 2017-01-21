package service

// Service is an interface that needs to be satisfied by defined services.
type Service interface {
	Start() error
	Stop() error
	// This should return a status and an error ([]string?) or have a status
	// object defined
	Status() ([]Status, error)
}

// Status structure holds information about the current service.
type Status struct {
	StartTime      string
	LivenessProbe  string
	ReadinessProbe string
}
