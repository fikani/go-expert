package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")
var ErrHandlerNotFound = errors.New("handler not found")
var ErrEventNotFound = errors.New("event not found")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) RegisterHandler(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	waitGroup := &sync.WaitGroup{}
	for _, handler := range ed.handlers[event.GetName()] {
		waitGroup.Add(1)
		go handler.Handle(event, waitGroup)
	}
	waitGroup.Wait()
	return nil
}

func (ed *EventDispatcher) ClearHandlers() error {
	ed.handlers = make(map[string][]EventHandlerInterface)
	return nil
}

func (ed *EventDispatcher) RemoveHandler(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; !ok {
		return ErrEventNotFound
	}

	for i, h := range ed.handlers[eventName] {
		if h == handler {
			ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
			return nil
		}
	}

	return ErrHandlerNotFound
}

func (ed *EventDispatcher) HasHandler(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; !ok {
		return false
	}

	for _, h := range ed.handlers[eventName] {
		if h == handler {
			return true
		}
	}

	return false
}
