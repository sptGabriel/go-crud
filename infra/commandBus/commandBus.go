package commandBus

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	InvalidHandlerERR = errors.New("this handler doenst exists on handlers map")
)

type Command interface {
}

type CommandHandler interface {
	Handle(Command) (result interface{}, err *CommandHandlerError)
}

type CommandBus struct {
	handlersMap map[string]CommandHandler
}

type CommandHandlerError struct {
	Error   error
	Code    int
	Message string
}

func NewCommandBus() CommandBus {
	return CommandBus{handlersMap: make(map[string]CommandHandler)}
}

func (cb *CommandBus) RegisterHandler(c Command, ch CommandHandler) error {
	cmdName := reflect.TypeOf(c).String()
	_, has := cb.handlersMap[cmdName]
	if has {
		return fmt.Errorf("the Command %s is already register", cmdName)
	}
	cb.handlersMap[cmdName] = ch
	return nil
}

func (cb CommandBus) Publish(c Command) (interface{}, *CommandHandlerError) {
	cmdName := reflect.TypeOf(c).String()
	ch, ok := cb.handlersMap[cmdName]
	if !ok {
		return nil, &CommandHandlerError{Error: InvalidHandlerERR, Code: 500}
	}
	return ch.Handle(c)
}
