package main

import "fmt"

//Begin Transactions
func (app *App) Begin() stateFn {
	app.TransactionState = CopyMap(app.Store)
	app.TransactionActive = true
	return app.Prompt
}

//Rollback Transactions
func (app *App) Rollback() stateFn {
	if app.checkIfTransactionsExist() {
		app.Store = app.TransactionState
	}
	app.resetTransactions()
	return app.Prompt
}

//Commit Transactions
func (app *App) Commit() stateFn {
	if app.checkIfTransactionsExist() {
		app.TransactionState = nil
	}
	app.resetTransactions()
	return app.Prompt
}

//Helper function to Reset Transactions
func (app *App) resetTransactions() {
	app.TransactionActive = false
	app.TransactionCount = 0
}

//Helper function to check if Transactions exist
func (app *App) checkIfTransactionsExist() bool {
	if app.TransactionCount == 0 {
		fmt.Println("NO TRANSACTION")
		return false
	}
	return true
}
