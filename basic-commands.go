package main

import "fmt"

//Set sets a value in the keystore for the provided key
func (app *App) Set(commands []string) stateFn {
	if len(commands) != 3 {
		return app.InvalidCommand("invalid set request")
	}
	if app.TransactionActive {
		app.TransactionCount += 1
	}
	key := commands[1]
	value := commands[2]
	app.Store[key] = value
	return app.Prompt
}

// UnSet  Deletes the key from the keystore
func (app *App) UnSet(commands []string) stateFn {
	if len(commands) != 2 {
		return app.InvalidCommand("incorrect number of variables")
	}
	if app.TransactionActive {
		app.TransactionCount += 1
	}
	key := commands[1]
	delete(app.Store, key)
	return app.Prompt
}

// NumEqualTo gets the numbers of values for the value for provided in command
func (app *App) NumEqualTo(commands []string) stateFn {
	if len(commands) != 2 {
		return app.InvalidCommand("incorrect number of variables")
	}
	val := commands[1]
	count := 0
	for _, value := range app.Store {
		if val == value {
			count += 1
		}
	}
	fmt.Println(count)
	return app.Prompt
}

// All  Prints all values in the keystore
func (app *App) All() stateFn {
	for key, value := range app.Store {
		fmt.Println("key:", key, "|value:", value)
	}
	return app.Prompt
}

//Get gets the value from the keystore based on the provided key
func (app *App) Get(commands []string) stateFn {
	if len(commands) != 2 {
		return app.InvalidCommand("Invalid get request")
	}
	key := commands[1]
	if app.Store[key] != "" {
		fmt.Println(app.Store[key])
	} else {
		fmt.Println("nil")
	}
	return app.Prompt
}

//End exits the program
func (app *App) End() stateFn {
	return nil
}
