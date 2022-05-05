package main

import "fmt"

func (app *App) Set(parsedCommands []string) stateFn {
	if len(parsedCommands) != 3 {
		return app.InvalidCommand("invalid set request")
	}
	if app.TransactionActive {
		app.TransactionCount += 1
	}
	key := parsedCommands[1]
	value := parsedCommands[2]
	app.Store[key] = value
	return app.Prompt
}

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

func (app *App) NumEqualTo(parsedCommands []string) stateFn {
	if len(parsedCommands) != 2 {
		return app.InvalidCommand("incorrect number of variables")
	}
	val := parsedCommands[1]
	count := 0
	for _, value := range app.Store {
		if val == value {
			count += 1
		}
	}
	fmt.Println(count)
	return app.Prompt
}

func (app *App) All() stateFn {
	for key, value := range app.Store {
		fmt.Println("key:", key, "|value:", value)
	}
	return app.Prompt
}

func (app *App) Get(parsedCommands []string) stateFn {
	if len(parsedCommands) != 2 {
		return app.InvalidCommand("Invalid get request")
	}
	key := parsedCommands[1]
	if app.Store[key] != "" {
		fmt.Println(app.Store[key])
	} else {
		fmt.Println("nil")
	}
	return app.Prompt
}

func (app *App) End() stateFn {
	return nil
}
