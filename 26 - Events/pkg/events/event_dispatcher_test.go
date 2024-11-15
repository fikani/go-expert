package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	name     string
	dateTime string
	payload  interface{}
}

func (te *TestEvent) GetName() string {
	return te.name
}

func (te *TestEvent) GetDateTime() time.Time {
	t, err := time.Parse(time.RFC3339, te.dateTime)
	if err != nil {
		return time.Now()
	}
	return t
}

func (te *TestEvent) GetPayload() interface{} {
	return te.payload
}

type TestEventHandler struct {
	ID string
}

func (teh *TestEventHandler) Handle(event EventInterface, waitGroup *sync.WaitGroup) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event      TestEvent
	event2     TestEvent
	handler    TestEventHandler
	handler2   TestEventHandler
	dispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.event = TestEvent{
		name:     "test",
		dateTime: time.Now().Format(time.RFC3339),
		payload:  "test",
	}
	suite.event2 = TestEvent{
		name:     "test2",
		dateTime: time.Now().Format(time.RFC3339),
		payload:  "test2",
	}
	suite.handler = TestEventHandler{}
	suite.handler2 = TestEventHandler{}
	suite.dispatcher = NewEventDispatcher()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_RegisterHandler() {
	err := suite.dispatcher.RegisterHandler(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.RegisterHandler(suite.event.GetName(), &suite.handler2)
	suite.NoError(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	suite.Equal(&suite.handler, suite.dispatcher.handlers[suite.event.GetName()][0])
	suite.Equal(&suite.handler2, suite.dispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_RegisterHandler_WhenRegisterSameHandler() {
	err := suite.dispatcher.RegisterHandler(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.RegisterHandler(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_ClearHandlers() {
	err := suite.dispatcher.RegisterHandler(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.RegisterHandler(suite.event2.GetName(), &suite.handler2)
	suite.NoError(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event2.GetName()]))

	err = suite.dispatcher.ClearHandlers()
	suite.NoError(err)
	suite.Equal(0, len(suite.dispatcher.handlers[suite.event.GetName()]))
	suite.Equal(0, len(suite.dispatcher.handlers[suite.event2.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_RemoveHandler() {
	err := suite.dispatcher.RegisterHandler(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.RemoveHandler(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Equal(0, len(suite.dispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_RemoveHandler_WhenMultipleHandlers() {
	handler3 := TestEventHandler{}
	err := suite.dispatcher.RegisterHandler(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	err = suite.dispatcher.RegisterHandler(suite.event.GetName(), &suite.handler2)
	suite.NoError(err)
	err = suite.dispatcher.RegisterHandler(suite.event.GetName(), &handler3)
	suite.NoError(err)
	suite.Equal(3, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.RemoveHandler(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	err = suite.dispatcher.RemoveHandler(suite.event.GetName(), &suite.handler2)
	suite.NoError(err)

	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))
	suite.Equal(&handler3, suite.dispatcher.handlers[suite.event.GetName()][0])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_RemoveHandler_WhenEventNotFound() {
	err := suite.dispatcher.RemoveHandler(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrEventNotFound, err)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_RemoveHandler_WhenHandlerNotFound() {
	err := suite.dispatcher.RegisterHandler(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.RemoveHandler(suite.event.GetName(), &suite.handler2)
	suite.Equal(ErrHandlerNotFound, err)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_HasHandler() {
	err := suite.dispatcher.RegisterHandler(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.True(suite.dispatcher.HasHandler(suite.event.GetName(), &suite.handler))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_HasHandler_WhenEventNotFound() {
	err := suite.dispatcher.RegisterHandler(suite.event.GetName(), &suite.handler)
	suite.NoError(err)

	suite.False(suite.dispatcher.HasHandler(suite.event2.GetName(), &suite.handler))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_HasHandler_WhenHandlerNotFound() {
	err := suite.dispatcher.RegisterHandler(suite.event.GetName(), &suite.handler)
	suite.NoError(err)

	suite.False(suite.dispatcher.HasHandler(suite.event.GetName(), &suite.handler2))
}

type MockHandler struct {
	mock.Mock
}

func (mh *MockHandler) Handle(event EventInterface, waitGroup *sync.WaitGroup) {
	waitGroup.Done()
	mh.Called(event)
}
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	// mock handler
	mockHandler := new(MockHandler)
	mockHandler.On("Handle", &suite.event).Return()

	mockHandler2 := new(MockHandler)
	mockHandler2.On("Handle", &suite.event).Return()

	err := suite.dispatcher.RegisterHandler(suite.event.GetName(), mockHandler)
	suite.NoError(err)

	err = suite.dispatcher.RegisterHandler(suite.event.GetName(), mockHandler2)
	suite.NoError(err)

	err = suite.dispatcher.Dispatch(&suite.event)
	suite.NoError(err)

	mockHandler.AssertExpectations(suite.T())
	mockHandler.AssertNumberOfCalls(suite.T(), "Handle", 1)
	mockHandler2.AssertExpectations(suite.T())
	mockHandler2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
