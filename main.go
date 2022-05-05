package main

import (
	"fmt"
	"strings"
)

const (
	GET        = "get"
	SET        = "set"
	UNSET      = "unset"
	NumEqualTo = "numequalto"
	END        = "end"
	ALL        = "all"
	BEGIN      = "begin"
	ROLLBACK   = "rollback"
	COMMIT     = "commit"
	HELP       = "help"
)

func main() {
	fmt.Println("Welcome to Simple KeyStore")
	//Setup Keystore
	var app = App{
		Store:             make(map[string]string),
		TransactionActive: false,
		TransactionCount:  0,
	}
	//Start State Machine
	for next := app.Help; next != nil; {
		next = next()
	}
	fmt.Println("Thank you for Using Simple DB")
}

// State Function type
type stateFn func() stateFn

//This Function takes all inputs and executes the correct funtion accordingly
func (app *App) InputHandler(parsedCommands []string) stateFn {

	if len(parsedCommands) == 0 {
		return app.InvalidCommand("no command provided")
	}

	switch strings.ToLower(parsedCommands[0]) {

	case SET:
		return app.Set(parsedCommands)
	case GET:
		return app.Get(parsedCommands)
	case UNSET:
		return app.UnSet(parsedCommands)
	case NumEqualTo:
		return app.NumEqualTo(parsedCommands)
	case END:
		return app.End
	case ALL:
		return app.All
	case HELP:
		return app.Help
	case BEGIN:
		return app.Begin
	case ROLLBACK:
		return app.Rollback
	case COMMIT:
		return app.Commit
	default:
		return app.InvalidCommand("invalid command")
	}

}

// App Store
type App struct {
	Store             DB
	TransactionState  DB
	TransactionCount  int
	TransactionActive bool
}

// DB Object Map
type DB map[string]string
