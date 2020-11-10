package queues

// Application represents a transaction application
type Application interface {
	Start() error
	Stop() error
}
