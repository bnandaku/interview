package main

import (
	"bufio"
	"fmt"
	"os"
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
	var app = App{
		Store:             make(map[string]string),
		TransactionActive: false,
		TransactionCount:  0,
	}
	app.InitDB()
}

type stateFn func() stateFn

func (app *App) InitDB() {
	for next := app.Help; next != nil; {
		next = next()
	}
	fmt.Println("Thank you for Using Simple DB")
}

func (app *App) NextStep(parsedCommands []string) stateFn {

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

func (app *App) InvalidCommand(message string) stateFn {
	fmt.Println("error: ", message)
	return app.Help
}

// StringPrompt asks for a string value using the label
func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func (app *App) Help() stateFn {

	fmt.Println("--- Basic Commands ---")
	fmt.Println("SET <name> <value> -- Sets a record with key <name> and Value <value>")
	fmt.Println("GET <name> -- Gets a record with key <name> prints nil if not found")
	fmt.Println("UNSET <name> -- Unsets a record with key <name>")
	fmt.Println("NUMEQUALTO <value> -- Prints number of records stored")
	fmt.Println("ALL -- Prints All Records")
	fmt.Println("HELP -- prints this message")
	fmt.Println("----Transaction Commands----")
	fmt.Println("Begin -- Begins a transaction session")
	fmt.Println("Rollback -- Rolls back the db before the transaction session")
	fmt.Println("commit -- commits the transactions to the store")
	fmt.Println("END -- Exits program")
	return app.Prompt

}

func (app *App) Prompt() stateFn {
	command := StringPrompt("db>")
	commands := strings.Split(command, " ")
	return app.NextStep(commands)
}

func CopyMap(cache Cache) Cache {

	var newCache Cache = make(map[string]string)
	for key, value := range cache {
		newCache[key] = value
	}
	return newCache
}

type App struct {
	Store             Cache
	TransactionState  Cache
	TransactionCount  int
	TransactionActive bool
}

type Cache map[string]string
