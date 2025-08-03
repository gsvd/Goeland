package store

import (
	"database/sql"
	"fmt"
)

type Account struct {
	ID       int
	JID      string
	Password string
}

func (s *Store) AddAccount(jid string, password string) (*Account, error) {
	res, err := s.db.Exec(`
		INSERT INTO accounts (jid, password)
		VALUES (?, ?)
	`, jid, password)
	if err != nil {
		return nil, fmt.Errorf("AddAccount: db.Exec: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("AddAccount: LastInsertId: %w", err)
	}

	return &Account{
		ID:       int(id),
		JID:      jid,
		Password: password,
	}, nil
}

func (s *Store) GetAllAccounts() ([]Account, error) {
	rows, err := s.db.Query(`SELECT id, jid, password FROM accounts`)
	if err != nil {
		return nil, fmt.Errorf("GetAllAccounts: db.Query: %w", err)
	}
	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		var a Account
		if err := rows.Scan(&a.ID, &a.JID, &a.Password); err != nil {
			return nil, fmt.Errorf("GetAllAccounts: rows.Scan: %w", err)
		}
		accounts = append(accounts, a)
	}
	return accounts, nil
}

func (s *Store) GetAccountByJID(jid string) (*Account, error) {
	row := s.db.QueryRow(`SELECT id, jid, password FROM accounts WHERE jid = ?`, jid)

	var a Account
	if err := row.Scan(&a.ID, &a.JID, &a.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("GetAccountByJID: no account found for JID %s", jid)
		}
		return nil, fmt.Errorf("GetAccountByJID: row.Scan: %w", err)
	}
	return &a, nil
}

func (s *Store) DeleteAccount(id int) error {
	_, err := s.db.Exec(`DELETE FROM accounts WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("DeleteAccount: db.Exec: %w", err)
	}
	return nil
}
