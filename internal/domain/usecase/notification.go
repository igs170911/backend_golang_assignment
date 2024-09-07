package usecase

// Notification interface
type Notification interface {
	Notify(address string, tx Transaction)
}
