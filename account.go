package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gsvd/goeland/internal/store"
	"github.com/xmppo/go-xmpp"
	"modernc.org/sqlite"
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

func (a *App) OpenAccount(jid string) error {
	account, err := a.store.GetAccountByJID(jid)
	if err != nil {
		return fmt.Errorf("GetAllAccounts: %w", err)
	}

	username := strings.Split(account.JID, "@")[0]
	host := strings.Split(account.JID, "@")[1]
	fmt.Printf("Connecting to %s@%s...\n", username, host)

	options := xmpp.Options{
		Host:          host,
		User:          account.JID,
		Password:      account.Password,
		Debug:         true,
		Session:       false,
		Status:        "xa",
		StatusMessage: "Available",

		NoTLS:    true,
		StartTLS: true,
	}

	talk, err := options.NewClient()
	if err != nil {
		fmt.Printf("OpenAccount: NewClient: %v\n", err)
		return fmt.Errorf("OpenAccount: NewClient: %w", err)
	}

	go func() {
		for {
			chat, err := talk.Recv()
			if err != nil {
				fmt.Printf("OpenAccount: Recv: %v\n", err)
			}
			switch v := chat.(type) {
			case xmpp.Chat:
				fmt.Println(v.Remote, v.Text)
			case xmpp.Presence:
				fmt.Println(v.From, v.Show)
			}
		}
	}()

	return nil
}
