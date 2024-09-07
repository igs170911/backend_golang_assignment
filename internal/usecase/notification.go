package usecase

import (
	"fmt"
	"parse_server/internal/domain/usecase"
)

// ConsoleNotification 實現了 Notification interface
type ConsoleNotification struct{}

func (n *ConsoleNotification) Notify(address string, tx usecase.Transaction) {
	fmt.Printf("Notification - New transaction for address %s: %+v\n", address, tx)
}

func MustNotification() usecase.Notification {
	return &ConsoleNotification{}
}
