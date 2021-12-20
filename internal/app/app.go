package app

import (
	"log"
)

type APP interface {
	Users
	Books
}

type app struct {
	log *log.Logger
}

func NewApp(log *log.Logger) APP {
	a := new(app)
	a.log = log

	return a
}
