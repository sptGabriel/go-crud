package commandBus

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ERRInvalidHandler = errors.New("this handler doenst exists on handlers map")
)

type Command interface {
}

type CommandHandler interface {
	Execute(Command) interface{}
}

type CommandBus struct {
	handlersMap map[reflect.Type]CommandHandler
}

type CommandHandlerError struct {
	Error   error
	Code    int
	Message string
}

func NewCommandBus() CommandBus {
	return CommandBus{handlersMap: make(map[reflect.Type]CommandHandler)}
}

func (cb *CommandBus) RegisterHandler(c reflect.Type, ch CommandHandler) error {
	_, has := cb.handlersMap[c]
	if has {
		return fmt.Errorf("Command is already register")
	}
	cb.handlersMap[c] = ch
	return nil
}

func (cb CommandBus) Publish(c Command) interface{} {
	cmdType := reflect.TypeOf(c)
	ch, ok := cb.handlersMap[cmdType]
	if !ok {
	}
	return ch.Execute(c)
}
