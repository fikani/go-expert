package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface, waitGroup *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Dispatch(event EventInterface) error
	RegisterHandler(eventName string, handler EventHandlerInterface) error
	RemoveHandler(eventName string, handler EventHandlerInterface) error
	HasHandler(eventName string, handler EventHandlerInterface) bool
	ClearHandlers() error
}
