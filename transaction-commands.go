package main

import "fmt"

func (app *App) Begin() stateFn {
	app.TransactionState = CopyMap(app.Store)
	app.TransactionActive = true
	return app.Prompt
}

func (app *App) Rollback() stateFn {
	if !app.checkIfTransactionsExist() {
		app.Store = app.TransactionState
	}
	app.resetTransactions()
	return app.Prompt
}

func (app *App) Commit() stateFn {
	if !app.checkIfTransactionsExist() {
		app.TransactionState = nil
	}
	app.resetTransactions()
	return app.Prompt
}

func (app *App) resetTransactions() {
	app.TransactionActive = false
	app.TransactionCount = 0
}

func (app *App) checkIfTransactionsExist() bool {
	if app.TransactionCount == 0 {
		fmt.Println("NO TRANSACTION")
		return false
	}
	return true
}
