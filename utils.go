package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// InvalidCommand - Catch all error message for all functions
func (app *App) InvalidCommand(message string) stateFn {
	fmt.Println("error: ", message)
	return app.Help
}

// StringPrompt Prompts the user to enter values
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

//Help provides helpful text for the user
func (app *App) Help() stateFn {
	fmt.Println("--- Basic Commands ---")
	fmt.Println("SET <name> <value> -- Sets a record with key <name> and Value <value>")
	fmt.Println("GET <name> -- Gets a record with key <name> prints nil if not found")
	fmt.Println("UNSET <name> -- Unsets a record with key <name>")
	fmt.Println("NUMEQUALTO <value> -- Prints number times a certain value appears in keystore")
	fmt.Println("ALL -- Prints All values in keystore")
	fmt.Println("----Transaction Commands----")
	fmt.Println("Begin -- Begins a transaction session")
	fmt.Println("Rollback -- Rolls back the keystore before the transaction session")
	fmt.Println("commit -- commits the transactions to the keystore")
	fmt.Println("--Other Commands--")
	fmt.Println("END -- Exits program")
	fmt.Println("HELP -- prints this message")
	fmt.Println("All commands are case insensitive but values stored in keystore are case sensitive")
	return app.Prompt

}

//Prompt Starts a prompt waiting for user input
func (app *App) Prompt() stateFn {
	command := StringPrompt("db>")
	commands := strings.Split(command, " ")
	return app.InputHandler(commands)
}

// CopyMap Helper function that copies the keystore and returns a new keystore.
func CopyMap(cache DB) DB {

	var newCache DB = make(map[string]string)
	for key, value := range cache {
		newCache[key] = value
	}
	return newCache
}
