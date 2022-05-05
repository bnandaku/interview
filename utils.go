package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	return app.InputHandler(commands)
}

func CopyMap(cache DB) DB {

	var newCache DB = make(map[string]string)
	for key, value := range cache {
		newCache[key] = value
	}
	return newCache
}
