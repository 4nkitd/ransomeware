package exception

import (
	"fmt"
	"time"
)

type Handler struct {
	e   Exception
	err error
}

var handler Handler

func (handler *Handler) Handle(err error) (e Exception) {

	handler.err = err

	defer handler.Recover()
	handler.ParseError()

	return handler.e
}

func (handler *Handler) ParseError() (e Exception) {
	handler.e = Exception{
		time:     time.Now().UTC(),
		message:  handler.err.Error(),
		location: "that",
	}
	return handler.e
}

func (handler *Handler) Recover() {
	recover()
	fmt.Println("Recovery Successfully from Exception")
}
