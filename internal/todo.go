package internal

import "errors"

type Todo struct {
	message string
	creator string
	state   int
}

func (t Todo) Message() string {
	msg := t.message

	if t.IsActive() {
		msg = msg + " AAAAAA"
	}

	return msg
}

func (t Todo) IsActive() bool {
	return t.state == 1
}

func NewTodo(msg, creator string, state int) (Todo, error) {
	if state > 10 {
		return Todo{}, errors.New("invalid state number")
	}

	return Todo{
		message: msg,
		creator: creator,
		state:   state,
	}, nil
}
