package commandBus

import (
	"fmt"
	"reflect"

	"github.com/sptGabriel/go-ddd/application/errors"
)

var (
	ErrInvalidHandler = fmt.Errorf("this handler doenst exists on handlers map")
)

type Command interface {
}

type CommandHandler interface {
	Execute(Command) (res interface{}, err error)
}

type CommandBus struct {
	handlersMap map[reflect.Type]CommandHandler
}

func NewCommandBus() CommandBus {
	return CommandBus{handlersMap: make(map[reflect.Type]CommandHandler)}
}

func (cb *CommandBus) RegisterHandler(c reflect.Type, ch CommandHandler) error {
	const op = "commandBus.registerHandler"
	_, has := cb.handlersMap[c]
	if has {
		return errors.E(op, fmt.Errorf("handler already exists"), 500)
	}
	cb.handlersMap[c] = ch
	return nil
}

func (cb CommandBus) Publish(c Command) (interface{}, error) {
	const op = "commandBus.publish"
	cmdType := reflect.TypeOf(c)
	ch, ok := cb.handlersMap[cmdType]
	if !ok {
		return nil, errors.E(op, ErrInvalidHandler, 500)
	}
	return ch.Execute(c)
}
