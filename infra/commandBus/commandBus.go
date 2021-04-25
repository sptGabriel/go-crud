package commandBus

import (
	"fmt"
	"reflect"
)

type Command interface {
}

type CommandHandler interface {
	Handle(Command) error
}

type CommandBus struct {
	handlersMap map[string]CommandHandler
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

func (cb CommandBus) Publish(c Command) (interface{}, error) {
	cmdName := reflect.TypeOf(c).String()
	ch, ok := cb.handlersMap[cmdName]
	if !ok {
		return nil, fmt.Errorf("there not any CommandHandler associate to Command %s", cmdName)
	}
	return ch.Handle(c), nil
}
