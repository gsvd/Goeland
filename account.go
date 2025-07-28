package main

import (
	"errors"
	"fmt"

	"modernc.org/sqlite"
	"github.com/gsvd/goeland/internal/store"
	lib "modernc.org/sqlite/lib"
)
func (a *App) AddAccount(jid string, password string) (*store.Account, error) {
	if jid == "" {
		return nil, NewAppError("ERR_EMPTY_JID")
	}
	if password == "" {
		return nil, NewAppError("ERR_PASSWORD_REQUIRED")
	}

	account, err := a.store.AddAccount(jid, password)
	if err != nil {
		var sqliteErr *sqlite.Error
		if errors.As(err, &sqliteErr) {
			if sqliteErr.Code() == lib.SQLITE_CONSTRAINT_UNIQUE {
				return nil, NewAppError("ERR_ACCOUNT_EXISTS")
			}
		}
		return nil, fmt.Errorf("AddAccount: %w", err)
	}

	return account, nil
}


func (a *App) GetAllAccounts() ([]store.Account, error) {
	accounts, err := a.store.GetAllAccounts()
	if err != nil {
		return nil, fmt.Errorf("GetAllAccounts: %w", err)
	}
	return accounts, nil
}
