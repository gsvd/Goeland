package store

import (
	"fmt"
)

type Account struct {
	ID       int
	JID      string
	Password string
}

func (s *Store) AddAccount(jid string, password string) error {
	_, err := s.db.Exec(`
		INSERT INTO accounts (jid, password)
		VALUES (?, ?)
	`, jid, password)
	if err != nil {
		return fmt.Errorf("AddAccount: db.Exec: %w", err)
	}
	return nil
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

func (s *Store) DeleteAccount(id int) error {
	_, err := s.db.Exec(`DELETE FROM accounts WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("DeleteAccount: db.Exec: %w", err)
	}
	return nil
}
