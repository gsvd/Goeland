package main

import (
	"errors"
	"fmt"

	"modernc.org/sqlite"
	lib "modernc.org/sqlite/lib"
)

func (a *App) AddAccount(jid string, password string) error {
	if jid == "" {
		return NewAppError("ERR_EMPTY_JID")
	}
	if password == "" {
		return NewAppError("ERR_PASSWORD_REQUIRED")
	}

	err := a.store.AddAccount(jid, password)
	if err != nil {
		var sqliteErr *sqlite.Error
		if errors.As(err, &sqliteErr) {
			if sqliteErr.Code() == lib.SQLITE_CONSTRAINT_UNIQUE {
				return NewAppError("ERR_ACCOUNT_EXISTS")
			}
		}
		return fmt.Errorf("AddAccount: %w", err)
	}

	return err
}

func (a *App) HasAnyAccount() (bool, error) {
	accounts, err := a.store.GetAllAccounts()
	if err != nil {
		return false, err
	}
	return len(accounts) > 0, nil
}
